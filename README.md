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

## To run `fireantelope` and `dfuseeos`

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

## To run `fireantelope` docker image and `dfuseeos`

1. Start `nodeos` on host machine
```
ultratest -s -n -D
```

2. Start `firehose-antelope` nodeos image

- Note: change docker image tag if needed. Change `$ULTRA_DIR` to point to your ultra directory.
```sh
docker run --rm -it --network="host" -v $ULTRA_DIR/firehose-antelope:/tmp -p 1066:1066 -p 10010:10010 -p 10012:10012 -p 10014:10014 eu.gcr.io/dfuse-302310/firehose-antelope:ci-test-nodeos3.2.4-2.0.3 bash
```

3. Start `fireantelope` inside the container
```
cd /tmp
./bin/run_local_test.sh docker
```

4. Modify `common-blocks-store-url` appropriately in `devel/standard/dfuse.yaml` and start `dfuseeos` from another console.
```
dfuseeos start -c ./devel/standard/dfuse.yaml
```
- Open block explorer at [localhost:8080](http:://localhost:8080) to validate.

5. (Optional) Run `SubstreamsCLI` from yet another console.
```
~/pinax-network/substreams/eosio.token$ substreams run -e localhost:10016 substreams.yaml map_transfers --start-block 2 --stop-block 20 --plaintext
```

### For WSL in Windows

- If not work try to update `firehose-data/reader/config/config.ini`, change `p2p-peer-address = 127.0.0.1:9876` to `p2p-peer-address = host.docker.internal:9876`

# Deployment

Todo: describe deployments for firehose and substreams

## License

[Apache 2.0](LICENSE)
