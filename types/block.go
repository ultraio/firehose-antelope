package types

import (
	"fmt"
	pbantelope "github.com/pinax-network/firehose-antelope/types/pb/sf/antelope/type/v1"
	"github.com/streamingfast/bstream"
	"github.com/streamingfast/logging"
	pbbstream "github.com/streamingfast/pbgo/sf/bstream/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var zlog, _ = logging.PackageLogger("types", "github.com/streamingfast/firehose-antelope/types")

func BlockFromProto(b *pbantelope.Block) (*bstream.Block, error) {

	content, err := proto.Marshal(b)
	if err != nil {
		zlog.Error("failed to marshal to binary", zap.Error(err), zap.Any("block", b))
		return nil, fmt.Errorf("unable to marshal to binary form: %s", err)
	}

	blk := &bstream.Block{
		Id:             b.ID(),
		Number:         b.Num(),
		PreviousId:     b.PreviousID(),
		Timestamp:      b.Time(),
		LibNum:         b.LIBNum(),
		PayloadKind:    pbbstream.Protocol_EOS,
		PayloadVersion: 1,
	}

	return bstream.GetBlockPayloadSetter(blk, content)
}
