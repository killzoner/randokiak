#!/bin/bash
set -e

(cd ../terraform/elasticsearch      && terraform init && terraform apply -auto-approve)
(cd ../terraform/pulsar             && terraform init && terraform apply -auto-approve)
(cd ../terraform/go                 && terraform init && terraform apply -auto-approve)
(cd ../terraform/webapp             && terraform init && terraform apply -auto-approve)
(cd ../terraform/ingress-lb         && terraform init && terraform apply -auto-approve)