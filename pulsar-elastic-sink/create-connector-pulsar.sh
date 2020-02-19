#!/bin/bash

# for local debug, replace `create` with `localrun` 
/pulsar/bin/pulsar-admin sinks create \
    --archive $PWD/connectors/pulsar-io-elastic-search-2.5.0.nar \
    --tenant public \
    --namespace default \
    --name elasticsearch-randokiak-sink \
    --sink-config-file $PWD/pulsar-es-sink.yml \
    --inputs randokiak-topic

# go localhost:9200/randokiak-index/_search to see if it works

# TODO : add script on a volume and execute command on a postStart lifecycle hook
# see https://stackoverflow.com/questions/44140593/how-to-run-command-after-initialization
# see https://etoews.github.io/blog/2017/07/29/inject-an-executable-script-into-a-container-in-kubernetes/ 