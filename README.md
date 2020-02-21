# randokiak

[![Build Status](https://travis-ci.org/killzoner/randokiak.svg?branch=dev)](https://travis-ci.org/killzoner/randokiak)

Random characters exposed in a very complicated but instructive way

## Getting started

### Why this project

I said "Nicolas, say the words !". And Nicolas said :

- Golang
- Protobuf / GRPC
- Elasticsearch
- Apache Pulsar
- Kubernetes (k8s)
- Terraform
- Typescript
- Docker / docker-compose
- ~~Helm~~
- ~~Ansible~~
- ~~Ruby~~

I put aside the last crossed items, and added Angular on top of Typescript (because, why not ! - I really like Angular), and came up with a project mixing all these together to get my hands on them.

The project is named `Randokiak`, a pun with the word 'random' and the name of a company.
The goal is to... well, generate random characters that you can see in the webapp once it has been swallowed by the whole machinery.

### Mandatory tools to install

- a bash compatible shell (bash on linux, git-bash or cygwin on windows)
- Docker
- docker-compose
- terraform
- kind, or minikube and virtualbox

On linux or mac, prefer `brew` (<https://docs.brew.sh/Homebrew-on-Linux)> to install most of them.

On windows, prefer `scoop` (<https://github.com/lukesampson/scoop>) or `chocolatey`.

Run `build-kind.sh` or `build-minikube.sh` depending on your local cluster install.
See `LOCAL_K8S_CLUSTER.md` for more infos about it.

You will be able to access all the services via the random NodePort attribued to each service (see in dashboard).

You can also access webapp service and rdkapi service via the defined urls "rdk.io/rdkapi" and "rdk.io/webapp" (exposed via an ingress, not a nodePort).
| Note: If you want to be able to access services via these predefined urls, you will have to define a line in you /etc/hosts file in the following way `<local_cluster_root_ip> rdk.io` |
| --- |

### Development build (add to previous)

- go 1.13.x minimum
- Node v12.x minimum
- make
- protobuf-compiler
- kubectl
