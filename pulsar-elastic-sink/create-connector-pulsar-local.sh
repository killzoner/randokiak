#!/bin/bash
set -e

# this script is run in a postStart lifecycle hook
# see https://stackoverflow.com/questions/44140593/how-to-run-command-after-initialization

# see https://pulsar.apache.org/docs/en/next/io-elasticsearch-sink/
# create config
echo '
configs:
    elasticSearchUrl: "http://localhost:9200"
    indexName: "randokiak-index"
' > pulsar-es-sink.yml

# wait for pulsar functions to be initialized
until $(curl --output /dev/null --silent --head --fail http://localhost:8080/admin/v2/functions/connectors); do
    printf '.'
    sleep 5
done

# list connectors
#curl --verbose http://localhost:8080/admin/v2/functions/connectors
# get elastic connector
NB_CONNECTORS=$(curl --silent http://localhost:8080/admin/v3/sinks/public/default/elasticsearch-randokiak-sink/status | grep "numInstances" | wc -l)
echo "$NB_CONNECTORS found";

# create connectors if not exist
if [ "$NB_CONNECTORS" = "0" ]; then
  # for local debug, replace `create` with `localrun` 
  /pulsar/bin/pulsar-admin sinks create \
    --archive $PWD/connectors/pulsar-io-elastic-search-2.5.0.nar \
    --tenant public \
    --namespace default \
    --name elasticsearch-randokiak-sink \
    --sink-config-file $PWD/pulsar-es-sink.yml \
    --inputs randokiak-topic
else
  echo "Connector already exists"
fi

# cleanup
rm pulsar-es-sink.yml