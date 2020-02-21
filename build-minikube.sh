#!/bin/bash
set -e

# get minikube docker daemon for building into it
export MINIKUBE_IP=$(minikube ip)
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://$MINIKUBE_IP:2376"
export DOCKER_CERT_PATH="$HOME/.minikube/certs"
export DOCKER_API_VERSION="1.23"

# build projects in minikube docker daemon
(cd scripts && ./build-projects.sh)

# add ingress to minikube
minikube addons enable ingress

# deploy apps using terraform
(cd scripts && ./terraform-deploy.sh)

# deploy dashboard and run proxy
minikube dashboard