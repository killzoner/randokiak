#!/bin/bash
set -e

(cd ../go && docker-compose build)
(cd ../pulsar-elastic-sink && docker-compose build)
(cd ../webapp && docker-compose build)