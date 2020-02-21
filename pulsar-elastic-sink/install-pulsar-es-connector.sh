#!/bin/bash
set -e

#wget https://archive.apache.org/dist/pulsar/pulsar-2.5.0/connectors/pulsar-io-elastic-search-2.5.0.nar
curl -O -J https://archive.apache.org/dist/pulsar/pulsar-2.5.0/connectors/pulsar-io-elastic-search-2.5.0.nar
mkdir -p connectors
mv pulsar-io-elastic-search-2.5.0.nar connectors
ls connectors