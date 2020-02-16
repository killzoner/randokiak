#!/bin/bash

wget https://archive.apache.org/dist/pulsar/pulsar-2.5.0/connectors/pulsar-io-elastic-search-2.5.0.nar
mkdir -p connectors
mv pulsar-io-elastic-search-2.5.0.nar connectors
ls connectors

# for local debug, replace `create` with `localrun` 
./pulsar-admin sinks create \
    --archive $PWD/connectors/pulsar-io-elastic-search-2.5.0.nar \
    --tenant public \
    --namespace default \
    --name elasticsearch-randokiak-sink \
    --sink-config-file $PWD/pulsar-es-sink.yml \
    --inputs randokiak-topic

# go localhost:9200/randokiak-index/_search to see if it works