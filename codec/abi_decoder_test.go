package codec

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
	"github.com/pinax-network/firehose-antelope/types/pb/sf/antelope/type/v1"
	"github.com/streamingfast/jsonpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func TestABIDecoder(t *testing.T) {
	type expectation struct {
		path string
		// If value is a hex string, it expects `rawData` to match it, otherwise, it expects `jsonData` to match it
		value string
	}

	type testData struct {
		name         string
		abiDumps     map[string]*eos.ABI
		blocks       []*pbantelope.Block
		expectations []expectation
	}

	in := func(blocks ...*pbantelope.Block) []*pbantelope.Block {
		return blocks
	}

	tokenABI1 := readABI(t, "token.1.abi.json")
	tokenABI2 := readABI(t, "token.2.abi.json")
	testABI1 := readABI(t, "test.1.abi.json")
	testABI2 := readABI(t, "test.2.abi.json")
	testABI3 := readABI(t, "test.3.abi.json")
	systemABI := readABI(t, "system.abi.json")

	softFailStatus := pbantelope.TransactionStatus_TRANSACTIONSTATUS_SOFTFAIL
	hardFailStatus := pbantelope.TransactionStatus_TRANSACTIONSTATUS_HARDFAIL

	tests := []testData{
		{
			name: "setabi and usage, same trace",
			blocks: in(testBlock(t, "00000002aa", "00000001aa",
				trxTrace(t,
					actionTraceSetABI(t, "test", 0, 1, testABI1),
					actionTrace(t, "test:test:act1", 1, 2, testABI1, `{"from":"test1"}`),
				),
			)),
			expectations: []expectation{
				{"block 0/trace 0/action 1", `{"from":"test1"}`},
			},
		},
		{
			name: "setabi and usage, same block, two traces",
			blocks: in(testBlock(t, "00000002aa", "00000001aa",
				trxTrace(t, actionTraceSetABI(t, "test", 0, 1, testABI1)),
				trxTrace(t, actionTrace(t, "test:test:act1", 0, 2, testABI1, `{"from":"test1"}`)),
			)),
			expectations: []expectation{
				{"block 0/trace 1/action 0", `{"from":"test1"}`},
			},
		},
		{
			name: "setabi and usage, two different blocks",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, actionTraceSetABI(t, "test", 0, 1, testABI1)),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, actionTrace(t, "test:test:act1", 0, 2, testABI1, `{"from":"test1"}`)),
				),
			),
			expectations: []expectation{
				{"block 1/trace 0/action 0", `{"from":"test1"}`},
			},
		},
		{
			name: "set multiple times, within same transaction, two different blocks",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t,
						actionTraceSetABI(t, "test", 0, 1, testABI1),
						actionTrace(t, "test:test:act1", 1, 2, testABI1, `{"from":"test1"}`),
						actionTraceSetABI(t, "test", 2, 3, testABI2),
						actionTrace(t, "test:test:act2", 2, 4, testABI2, `{"to":20}`),
						actionTraceSetABI(t, "test", 4, 5, testABI3),
					),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, actionTrace(t, "test:test:act3", 0, 6, testABI3, `{"quantity":"1.0 EOS"}`)),
				),
			),
			expectations: []expectation{
				{"block 0/trace 0/action 1", `{"from":"test1"}`},
				{"block 0/trace 0/action 3", `{"to":20}`},
				{"block 1/trace 0/action 0", `{"quantity":"1.0 EOS"}`},
			},
		},
		{
			name: "set multiple times, across transactions, two different blocks",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, actionTraceSetABI(t, "test", 0, 1, testABI1)),
					trxTrace(t, actionTrace(t, "test:test:act1", 0, 2, testABI1, `{"from":"test1"}`)),
					trxTrace(t, actionTraceSetABI(t, "test", 0, 3, testABI2)),
					trxTrace(t, actionTrace(t, "test:test:act2", 0, 4, testABI2, `{"to":20}`)),
					trxTrace(t, actionTraceSetABI(t, "test", 0, 5, testABI3)),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, actionTrace(t, "test:test:act3", 0, 6, testABI3, `{"quantity":"1.0 EOS"}`)),
				),
			),
			expectations: []expectation{
				{"block 0/trace 1/action 0", `{"from":"test1"}`},
				{"block 0/trace 3/action 0", `{"to":20}`},
				{"block 1/trace 0/action 0", `{"quantity":"1.0 EOS"}`},
			},
		},
		{
			name: "fork multiple block",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, actionTraceSetABI(t, "test", 0, 1, testABI1)),
					trxTrace(t, actionTraceSetABI(t, "token", 0, 2, tokenABI1)),
				),
				testBlock(t, "00000002bb", "00000001aa",
					trxTrace(t, actionTrace(t, "test:test:act1", 0, 3, testABI1, `{"from":"test1"}`)),
					trxTrace(t, actionTraceSetABI(t, "test", 0, 4, testABI2)),
					trxTrace(t, actionTrace(t, "test:test:act2", 0, 5, testABI2, `{"to":20}`)),
				),
				testBlock(t, "00000003bb", "00000002bb",
					trxTrace(t, actionTrace(t, "test:test:act2", 0, 6, testABI2, `{"to":20}`)),
					trxTrace(t, actionTraceSetABI(t, "token", 0, 7, tokenABI2)),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, actionTrace(t, "test:test:act1", 0, 3, testABI1, `{"from":"test1"}`)),
					trxTrace(t, actionTrace(t, "token:token:transfer", 0, 4, tokenABI1, `{"to":"transfer3"}`)),
				),
			),
			expectations: []expectation{
				{"block 1/trace 0/action 0", `{"from":"test1"}`},
				{"block 1/trace 2/action 0", `{"to":20}`},
				{"block 2/trace 0/action 0", `{"to":20}`},
				{"block 3/trace 0/action 0", `{"from":"test1"}`},
				{"block 3/trace 1/action 0", `{"to":"transfer3"}`},
			},
		},
		{
			name: "fail transaction, does not save ABI",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, hardFailStatus, actionTraceSetABI(t, "test", 0, 1, testABI1)),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, actionTrace(t, "test:test:act1", 0, 2, testABI1, `{"from":"test1"}`)),
				),
			),
			expectations: []expectation{
				{"block 1/trace 0/action 0", `000000008090b1ca`},
			},
		},
		{
			name: "fail transaction, still works from failed transaction but does not record ABI",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, actionTraceSetABI(t, "test", 0, 1, testABI1)),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, hardFailStatus,
						actionTrace(t, "test:test:act1", 0, 2, testABI1, `{"from":"test1"}`),
						actionTraceSetABI(t, "test", 1, 3, testABI2),
						actionTrace(t, "test:test:act2", 2, 4, testABI2, `{"to":1}`),
						actionTrace(t, "test:test:act2", 3, 5, testABI2, `{"to":2}`),
						actionTraceSetABI(t, "test", 4, 6, testABI3),
						actionTraceFail(t, "test:test:act3", 5, testABI3, `{"quantity":"1.0000 EOS"}`),
					),
				),
				testBlock(t, "00000004aa", "00000003aa",
					trxTrace(t,
						actionTrace(t, "test:test:act1", 0, 2, testABI1, `{"from":"test3"}`),
						// Let's assume there is a bunch of transaction in-between, so we test that no recording actually occurred!
						actionTrace(t, "test:test:act1", 1, 7, testABI1, `{"from":"test4"}`),
					),
				),
			),
			expectations: []expectation{
				{"block 1/trace 0/action 0", `{"from":"test1"}`},
				{"block 1/trace 0/action 2", `{"to":1}`},
				{"block 1/trace 0/action 3", `{"to":2}`},
				{"block 1/trace 0/action 5", `{"quantity":"1.0000 EOS"}`},
				{"block 2/trace 0/action 0", `{"from":"test3"}`},
				{"block 2/trace 0/action 1", `{"from":"test4"}`},
			},
		},

		{
			name: "soft_fail onerror, still records ABI",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, softFailStatus,
						actionTrace(t, "eosio:eosio:onerror", 0, 1, nil, ""),
						actionTraceSetABI(t, "test", 1, 2, testABI2),
						actionTrace(t, "test:test:act2", 2, 3, testABI2, `{"to":1}`),
						actionTraceSetABI(t, "test", 3, 4, testABI3),
					),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, actionTrace(t, "test:test:act3", 0, 5, testABI3, `{"quantity":"1.0000 EOS"}`)),
				),
			),
			expectations: []expectation{
				{"block 0/trace 0/action 2", `{"to":1}`},
				{"block 1/trace 0/action 0", `{"quantity":"1.0000 EOS"}`},
			},
		},

		{
			name: "soft_fail, with abi dumps, single action global sequence 0, still records ABI",
			abiDumps: map[string]*eos.ABI{
				"eosio.token": tokenABI2,
			},
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, actionTraceSetABI(t, "eosio.token", 0, 1, tokenABI2)),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, softFailStatus,
						actionTraceFail(t, "eosio.token:eosio.token:transfer", 0, tokenABI2, `{"from":"bitfinexcw11","memo":"Simple test","quantity":"1.0000 EOS","to":"bitfinexcw12"}`),
					),
				),
			),
			expectations: []expectation{
				{"block 1/trace 0/action 0", `{"from":"bitfinexcw11","memo":"Simple test","quantity":"1.0000 EOS","to":"bitfinexcw12"}`},
			},
		},

		{
			name: "hard_fail onerror, still works from failed transaction but does not record ABI",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, hardFailStatus,
						actionTrace(t, "eosio:eosio:onerror", 0, 1, nil, ""),
						actionTraceSetABI(t, "test", 1, 2, testABI2),
						actionTrace(t, "test:test:act2", 2, 3, testABI2, `{"to":1}`),
						actionTraceSetABI(t, "test", 3, 4, testABI3),
						actionTraceFail(t, "any:any:any", 4, nil, ""),
					),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, actionTrace(t, "test:test:act3", 0, 1, testABI3, `{"quantity":"1.0000 EOS"}`)),
					// Let's assume there is a bunch of transaction in-between, so we test that no recording actually occurred!
					trxTrace(t, actionTrace(t, "test:test:act3", 0, 8, testABI3, `{"quantity":"2.0000 EOS"}`)),
				),
			),
			expectations: []expectation{
				{"block 0/trace 0/action 2", `{"to":1}`},
				{"block 1/trace 0/action 0", `102700000000000004454f5300000000`},
				{"block 1/trace 1/action 0", `204e00000000000004454f5300000000`},
			},
		},

		{
			name: "dtrx ops are correctly decoded",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t,
						actionTraceSetABI(t, "test", 0, 1, testABI1),
						actionTraceSetABI(t, "token", 1, 2, tokenABI1),
						actionTrace(t, "test:test:act1", 2, 3, testABI1, `{"from":"block1"}`),

						// A dtrx op created by action index 2
						dtrxOp(t, 2, "create", signedTrx(t,
							cfaAction(t, "token:transfer", tokenABI1, `{"to":"someone"}`),
							action(t, "test:act1", testABI1, `{"from":"inner1"}`),
						)),
					),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t,
						actionTrace(t, "test:test:act1", 0, 4, testABI1, `{"from":"block2"}`),

						// A dtrx op created by action index 0
						dtrxOp(t, 0, "create", signedTrx(t,
							cfaAction(t, "token:transfer", tokenABI1, `{"to":"somelse"}`),
							action(t, "test:act1", testABI1, `{"from":"inner2"}`),
						)),
					),
				),
				testBlock(t, "00000004aa", "00000003aa",
					trxTrace(t, dtrxOp(t, 0, "push_create", signedTrx(t, action(t, "test:act1", testABI1, `{"from":"push1"}`)))),
				),
			),
			expectations: []expectation{
				{"block 0/trace 0/action 2", `{"from":"block1"}`},
				{"block 0/trace 0/dtrxOp 0/action 0", `{"from":"inner1"}`},
				{"block 0/trace 0/dtrxOp 0/cfaAction 0", `{"to":"someone"}`},

				{"block 1/trace 0/dtrxOp 0/action 0", `{"from":"inner2"}`},
				{"block 1/trace 0/dtrxOp 0/cfaAction 0", `{"to":"somelse"}`},

				{"block 2/trace 0/dtrxOp 0/action 0", `{"from":"push1"}`},
			},
		},

		{
			name: "trx ops are correctly decoded",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t,
						actionTraceSetABI(t, "test", 0, 1, testABI1),
						actionTraceSetABI(t, "token", 1, 2, tokenABI1),
						actionTraceSetABI(t, "eosio", 2, 3, systemABI),
					),
				),
				testBlock(t, "00000003aa", "00000002aa",
					trxTrace(t, softFailStatus,
						actionTrace(t, "test:test:act1", 0, 4, testABI1, `{"from":"block2"}`),
					),
					trxTrace(t, hardFailStatus,
						actionTrace(t, "eosio:eosio:onerror", 0, 5, systemABI, `{"trx_id":"abc"}`),
					),

					trxOp(t, signedTrx(t,
						action(t, "eosio:onblock", systemABI, `{"id":"00000003aa"}`),
						cfaAction(t, "test:act1", testABI1, `{"from":"block3"}`),
					)),
					trxOp(t, signedTrx(t,
						action(t, "eosio:onerror", systemABI, `{"trx_id":"abc"}`),
						cfaAction(t, "token:transfer", tokenABI1, `{"to":"someone"}`),
					)),
				),
			),
			expectations: []expectation{
				{"block 1/trace 0/action 0", `{"from":"block2"}`},
				{"block 1/trace 1/action 0", `{"trx_id":"abc"}`},

				{"block 1/trxOp 0/action 0", `{"id":"00000003aa"}`},
				{"block 1/trxOp 0/cfaAction 0", `{"from":"block3"}`},
				{"block 1/trxOp 1/action 0", `{"trx_id":"abc"}`},
				{"block 1/trxOp 1/cfaAction 0", `{"to":"someone"}`},
			},
		},

		{
			name: "native eosio:transfer correctly decoded",
			blocks: in(
				testBlock(t, "00000002aa", "00000001aa",
					trxTrace(t, actionTraceSetABI(t, "eosio.token", 0, 1, tokenABI2)),
					trxTrace(t, actionTrace(t, "eosio.token:eosio.token:transfer", 0, 2, tokenABI2, `{"from":"eosio","to":"token","quantity":"1.0000 EOS","memo":""}`)),
					trxTrace(t, actionTrace(t, "eosio.token:eosio.token:transfer", 0, 3, tokenABI2, `{"from":"eosio","to":"token","quantity":"1.0000 EOS","memo":"With memo"}`)),
				),
			),
			expectations: []expectation{
				{"block 0/trace 1/action 0", `{"from":"eosio","to":"token","quantity":"1.0000 EOS","memo":""}`},
				{"block 0/trace 2/action 0", `{"from":"eosio","to":"token","quantity":"1.0000 EOS","memo":"With memo"}`},
			},
		},
		// TODO: Add those tests
		//        - ensures "hard-coded" system methods like `setabi`, `setcode` always work?
	}

	toString := func(in proto.Message) string {

		res, err := proto.Marshal(in)
		require.NoError(t, err)
		return string(res)

		// todo legacy code, check if properly replaced
		//out, err := (&jsonpb.Marshaler{}).MarshalToString(in)
		//require.NoError(t, err)
		//
		//return out
	}

	hexRegex := regexp.MustCompile(`^[0-9a-fA-F]+$`)
	actionTraceRegex := regexp.MustCompile(`^block (\d+)/trace (\d+)/action (\d+)$`)
	dtrxOpRegex := regexp.MustCompile(`^block (\d+)/trace (\d+)/dtrxOp (\d+)/(action|cfaAction) (\d+)$`)
	trxOpRegex := regexp.MustCompile(`^block (\d+)/trxOp (\d+)/(action|cfaAction) (\d+)$`)

	toInt := func(in string) int {
		out, err := strconv.ParseInt(in, 10, 32)
		require.NoError(t, err)

		return int(out)
	}

	extractTrace := func(testData *testData, regexMatch []string) (block *pbantelope.Block, trace *pbantelope.TransactionTrace) {
		block = testData.blocks[toInt(regexMatch[1])]
		trace = block.UnfilteredTransactionTraces[toInt(regexMatch[2])]
		return
	}

	assertMatchAction := func(expected string, action *pbantelope.Action) {
		if hexRegex.MatchString(expected) {
			require.Equal(t, expected, hex.EncodeToString(action.RawData), toString(action))
			require.Empty(t, action.JsonData, "JsonData should be empty\n%s", toString(action))
		} else {
			require.NotEmpty(t, action.RawData, "RawData should still be populated\n%s", toString(action))
			require.NotEmpty(t, action.JsonData, "JsonData should not be empty\n%s", toString(action))
			assert.JSONEq(t, expected, action.JsonData)
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			decoder := newABIDecoder()

			for contract, abi := range test.abiDumps {
				abiBinary, err := eos.MarshalBinary(abi)
				require.NoError(t, err)

				_ = decoder.addInitialABI(contract, base64.RawStdEncoding.EncodeToString(abiBinary))
			}

			for _, block := range test.blocks {
				maybePrintBlock(t, block)

				err := decoder.startBlock(block.Num())
				require.NoError(t, err)

				for _, trxTrace := range block.UnfilteredTransactionTraces {
					err := decoder.processTransaction(trxTrace)
					require.NoError(t, err)
				}

				// This should wait for all decoding in the block to terminate
				err = decoder.endBlock(block)
				require.NoError(t, err)
			}

			for _, expect := range test.expectations {
				var match []string

				if match = fullMatchRegex(actionTraceRegex, expect.path); match != nil {
					_, trace := extractTrace(&test, match)
					assertMatchAction(expect.value, trace.ActionTraces[toInt(match[3])].Action)
					continue
				}

				if match = fullMatchRegex(dtrxOpRegex, expect.path); match != nil {
					_, trace := extractTrace(&test, match)
					dtrxOp := trace.DtrxOps[toInt(match[3])]

					if match[4] == "cfaAction" {
						assertMatchAction(expect.value, dtrxOp.Transaction.Transaction.ContextFreeActions[toInt(match[5])])
					} else if match[4] == "action" {
						assertMatchAction(expect.value, dtrxOp.Transaction.Transaction.Actions[toInt(match[5])])
					}
					continue
				}

				if match = fullMatchRegex(trxOpRegex, expect.path); match != nil {
					block := test.blocks[toInt(match[1])]
					trxOp := block.UnfilteredImplicitTransactionOps[toInt(match[2])]

					if match[3] == "cfaAction" {
						assertMatchAction(expect.value, trxOp.Transaction.Transaction.ContextFreeActions[toInt(match[4])])
					} else if match[3] == "action" {
						assertMatchAction(expect.value, trxOp.Transaction.Transaction.Actions[toInt(match[4])])
					}
					continue
				}

				assert.Fail(t, "Unable to assert unknown expectation", "Expecation path %q not matching any assertion regex", expect.path)
			}
		})
	}
}

func fullMatchRegex(regex *regexp.Regexp, content string) []string {
	match := regex.FindAllStringSubmatch(content, -1)
	if match == nil {
		return nil
	}

	return match[0]
}

func testBlock(t *testing.T, blkID string, previousBlkID string, elements ...interface{}) *pbantelope.Block {
	pbblock := &pbantelope.Block{
		Id:     blkID,
		Number: eos.BlockNum(blkID),
	}

	blockTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05.5Z")
	require.NoError(t, err)

	pbblock.DposIrreversibleBlocknum = pbblock.Number - 1
	pbblock.Header = &pbantelope.BlockHeader{
		Previous:  previousBlkID,
		Producer:  "tester",
		Timestamp: timestamppb.New(blockTime),
	}

	for _, element := range elements {
		switch v := element.(type) {
		case *pbantelope.TransactionTrace:
			pbblock.UnfilteredTransactionTraceCount++
			pbblock.UnfilteredTransactionTraces = append(pbblock.UnfilteredTransactionTraces, v)
		case *pbantelope.TrxOp:
			pbblock.UnfilteredImplicitTransactionOps = append(pbblock.UnfilteredImplicitTransactionOps, v)
		}
	}

	return pbblock
}

func trxTrace(t *testing.T, elements ...interface{}) *pbantelope.TransactionTrace {
	trace := &pbantelope.TransactionTrace{
		Receipt: &pbantelope.TransactionReceiptHeader{
			Status: pbantelope.TransactionStatus_TRANSACTIONSTATUS_EXECUTED,
		},
	}

	for _, element := range elements {
		switch v := element.(type) {
		case *pbantelope.ActionTrace:
			trace.ActionTraces = append(trace.ActionTraces, v)
		case *pbantelope.DBOp:
			trace.DbOps = append(trace.DbOps, v)
		case *pbantelope.DTrxOp:
			trace.DtrxOps = append(trace.DtrxOps, v)
		case *pbantelope.TableOp:
			trace.TableOps = append(trace.TableOps, v)
		case pbantelope.TransactionStatus:
			trace.Receipt.Status = v
		}
	}

	return trace
}

func signedTrx(t *testing.T, elements ...interface{}) *pbantelope.SignedTransaction {
	signedTrx := &pbantelope.SignedTransaction{}
	signedTrx.Transaction = trx(t, elements...)

	return signedTrx
}

type ContextFreeAction *pbantelope.Action

func trx(t *testing.T, elements ...interface{}) *pbantelope.Transaction {
	trx := &pbantelope.Transaction{}

	for _, element := range elements {
		switch v := element.(type) {
		case *pbantelope.Action:
			trx.Actions = append(trx.Actions, v)
		case ContextFreeAction:
			trx.ContextFreeActions = append(trx.ContextFreeActions, (*pbantelope.Action)(v))
		}
	}

	return trx
}

func actionTrace(t *testing.T, tripletName string, executionIndex uint32, globalSequence uint64, abi *eos.ABI, data string) *pbantelope.ActionTrace {
	parts := strings.Split(tripletName, ":")
	receiver := parts[0]
	account := parts[1]
	actionName := parts[2]

	return &pbantelope.ActionTrace{
		ExecutionIndex: executionIndex,
		Receiver:       receiver,
		Receipt: &pbantelope.ActionReceipt{
			Receiver:       receiver,
			GlobalSequence: globalSequence,
		},
		Action: action(t, account+":"+actionName, abi, data),
	}
}

func actionTraceFail(t *testing.T, tripletName string, executionIndex uint32, abi *eos.ABI, data string) *pbantelope.ActionTrace {
	out := actionTrace(t, tripletName, executionIndex, 0, abi, data)
	out.Receipt = nil

	return out
}

func actionTraceSetABI(t *testing.T, account string, executionIndex uint32, globalSequence uint64, abi *eos.ABI) *pbantelope.ActionTrace {
	abiData, err := eos.MarshalBinary(abi)
	require.NoError(t, err)

	setABI := &system.SetABI{Account: eos.AccountName(account), ABI: eos.HexBytes(abiData)}
	rawData, err := eos.MarshalBinary(setABI)
	require.NoError(t, err)

	return &pbantelope.ActionTrace{
		ExecutionIndex: executionIndex,
		Receiver:       "eosio",
		Receipt: &pbantelope.ActionReceipt{
			Receiver:       "eosio",
			GlobalSequence: globalSequence,
		},
		Action: &pbantelope.Action{
			Account: "eosio",
			Name:    "setabi",
			RawData: rawData,
		},
	}
}

func cfaAction(t *testing.T, pairName string, abi *eos.ABI, data string) ContextFreeAction {
	return ContextFreeAction(action(t, pairName, abi, data))
}

func action(t *testing.T, pairName string, abi *eos.ABI, data string) *pbantelope.Action {
	parts := strings.Split(pairName, ":")
	account := parts[0]
	actionName := parts[1]

	var rawData []byte
	if abi != nil && data != "" {
		var err error
		rawData, err = abi.EncodeAction(eos.ActionName(actionName), []byte(data))
		require.NoError(t, err)
	}

	return &pbantelope.Action{
		Account: account,
		Name:    actionName,
		RawData: rawData,
	}
}

func trxOp(t *testing.T, signedTrx *pbantelope.SignedTransaction) *pbantelope.TrxOp {
	op := &pbantelope.TrxOp{
		Transaction: signedTrx,
	}

	return op
}

func dtrxOp(t *testing.T, actionIndex uint32, operation string, signedTrx *pbantelope.SignedTransaction) *pbantelope.DTrxOp {
	opName := pbantelope.DTrxOp_Operation_value["OPERATION_"+strings.ToUpper(operation)]

	op := &pbantelope.DTrxOp{
		Operation:   pbantelope.DTrxOp_Operation(opName),
		ActionIndex: actionIndex,
		Transaction: signedTrx,
	}

	return op
}

func maybePrintBlock(t *testing.T, block *pbantelope.Block) {
	if os.Getenv("DEBUG") == "" && os.Getenv("TRACE") != "true" {
		return
	}

	marshaler := &jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(block)
	require.NoError(t, err)

	// We re-normalize to a plain map[string]interface{} so it's printed as JSON and not a proto default String implementation
	normalizedOut := map[string]interface{}{}
	require.NoError(t, json.Unmarshal([]byte(out), &normalizedOut))

	zlog.Debug("processing test block", zap.Any("block", normalizedOut))
}
