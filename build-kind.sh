#!/bin/bash
set -e

# build projects
(cd scripts && ./build-projects.sh)

# load images in kind
(cd scripts && ./kind-load-images.sh)

# deploy ingress
(cd k8s/ingress && ./install-ingress.sh)

# deploy apps using terraform
(cd scripts && ./terraform-deploy.sh)

# deploy dashboard and run proxy
(cd k8s/dashboard && ./install-dashboard.sh)

