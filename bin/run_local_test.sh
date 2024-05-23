#!/usr/bin/env bash

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd .. && pwd )"
BIN=fireantelope
CONFIG=./devel/standard/standard.yaml

rm -rf $ROOT/firehose-data/localdata
rm -rf $ROOT/firehose-data/storage
rm -rf $ROOT/firehose-data/app.log.json
rm -rf $ROOT/firehose-data/reader/data
rm -rf $ROOT/firehose-data/reader/work
rm -rf $ROOT/firehose-data/reader/config/protocol_features

cd $ROOT
if [ -f "$ROOT/fireantelope" ]
then
    BIN=./fireantelope
fi

if [ "$1" == "docker" ]
then
    CONFIG=./devel/standard/standard-docker.yaml
fi

$BIN start -c $CONFIG