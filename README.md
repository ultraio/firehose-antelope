# Firehose on Antelope
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is Streamingfast's [Firehose](https://firehose.streamingfast.io) implementation adapted for Antelope chains, 
effectively enabling [Substreams](https://substreams.streamingfast.io) on all Antelope chains. 

# Usage

## Release

Use the `./bin/release.sh` Bash script to perform a new release. It will ask you questions
as well as driving all the required commands, performing the necessary operation automatically.
The Bash script runs in dry-mode by default, so you can check first that everything is all right.

Releases are performed using [goreleaser](https://goreleaser.com/).

## To run `fireantelope` and `dfuseoes`

1. Start `nodeos`.
```
~/ultra/firehose-antelope$ ultratest -s -n -D
```
2. Modify `reader-node-arguments` in `devel/standard/standard.yaml` and start `fireantelope`.
```
~/ultra/firehose-antelope$ ./fireantelope start -c ./devel/standard/standard.yaml
```

3. Modify `common-blocks-store-url` appropriately in `devel/standard/dfuse.yaml` and start `dfuseeos` from another console.
```
~/ultra/firehose-antelope$ dfuseeos start -c ./devel/standard/dfuse.yaml
```
4. Run `SubstreamsCLI` from yet another console.
```
~/pinax-network/substreams/eosio.token$ substreams run -e localhost:10016 substreams.yaml map_transfers --start-block 2 --stop-block 20 --plaintext
```
5. To restart from genesis, you need to remove several files/directories.
```
~/ultra/firehose-antelope$ rm -rf firehose-data/localdata/ ; rm -rf firehose-data/reader/config/protocol_features/ ; rm -rf firehose-data/reader/data/ ; rm -rf firehose-data/reader/work/ ; rm -rf firehose-data/storage/ ; rm firehose-data/app.log.json
~/ultra/firehose-antelope$ rm -rf dfuse-data/
```

# Deployment

Todo: describe deployments for firehose and substreams

## License

[Apache 2.0](LICENSE)
