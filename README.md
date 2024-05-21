# Firehose on Antelope

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is the Antelope chain-specific implementation part of firehose-core and enables both
[Firehose](https://firehose.streamingfast.io/introduction/firehose-overview)
and [Substreams](https://substreams.streamingfast.io) on Antelope chains with native blocks.

## For Developers

To get started with Firehose and Substreams, you need to sign up on https://app.pinax.network to get yourself an api
key. You'll also find quickstarts there to get you started and all of our available endpoints (we currently provide both
Firehose and Substreams endpoints for EOS, WAX and Telos, as well as different testnets).

For connecting to **Firehose** endpoints, you'll need the Protobufs which are published on
[buf.build](https://buf.build/pinax/firehose-antelope/docs/main). Some Golang example code on how to set up a Firehose
client can be found [here](https://github.com/pinax-network/firehose-examples-go).

To **consume** Antelope Substreams, please have a look at the
[documentation](https://substreams.streamingfast.io/documentation/consume). You can also find Substreams to deploy in
our Substreams repository [here](https://github.com/pinax-network/substreams) and on
[substreams.dev](https://substreams.dev).

To **develop** Antelope Substreams, have a look at
the [documentation](https://substreams.streamingfast.io/documentation/develop) here and at the Pinax SDK for Antelope
Substreams which can be found [here](https://github.com/pinax-network/substreams-antelope).

A collection of resources around Substreams can also be found
on [Awesome Substreams](https://github.com/pinax-network/awesome-substreams).

### Subgraphs

Although The Graph does not officially support any Antelope chains yet, it's possible to write Subgraphs based on
Substreams and set up your own Graph node. You can find an
example [here](https://github.com/pinax-network/eosio.token-subgraph) on how to achieve this.

## For Operators

Please have a look at the documentation [here](https://firehose.streamingfast.io) on how to set up your own Firehose &
Substreams stack. Note that indexing larger Antelope chains such as EOS or WAX requires parallel processing of the chain
and a lot of resources to have the indexing done in a reasonable time frame.

### EOS EVM

This implementation provides native Antelope blocks, including all Antelope specific block data. In case you are looking
for operating Firehose & Substreams for EOS EVM, please have a look at
the [firehose-ethereum](https://github.com/streamingfast/firehose-ethereum) repository; it provides a generic evm poller
to poll the EVM blocks from an RPC node.

## Support

In case of any questions around the Pinax endpoints or technology, feel free to hit us on
our [Discord server](https://discord.gg/pinax). For more generic questions around Substreams, you might also find the
[Streamingfast Discord server](https://discord.gg/jZwqxJAvRs) useful.

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

## Deployment

Todo: describe deployments for firehose and substreams

## License

[Apache 2.0](LICENSE)
