#!/bin/bash
set -e

kind load docker-image randokiak/randokiak-api:0.0.1
kind load docker-image randokiak/randokiak-gen:0.0.1
kind load docker-image randokiak/pulsar-elastic-sink:0.0.1
kind load docker-image randokiak/webapp:0.0.1
