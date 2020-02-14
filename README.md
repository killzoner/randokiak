# randokiak

[![Build Status](https://travis-ci.org/killzoner/randokiak.svg?branch=dev)](https://travis-ci.org/killzoner/randokiak)

Random characters exposed in a very complicated but instructive way

## Getting started

On windows, install `scoop` (<https://github.com/lukesampson/scoop),> then all the necessary elements:

- `scoop bucket add extras`
- `scoop portable-virtualbox terraform helm minikube kubectl`

- `minikube start`
- `minikube addons enable ingress` #enable ingress for nginx example
See <https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/>
- `minikube dashboard`
- `terraform plan`
- `terraform apply`

You can replace minikube with `kind` if you prefer to leverage existing docker installation (<https://github.com/kubernetes-sigs/kind)>

Also see:

- <https://kubernetes.io/fr/docs/setup/learning-environment/minikube/>
- <https://www.terraform.io/docs/providers/kubernetes/guides/getting-started.html>
- <https://learn.hashicorp.com/terraform/getting-started/intro>
- <https://itmecho.com/managing-kubernetes-with-terraform/>
- <https://www.terraform.io/docs/configuration/>

Docker and proxy conf, make sure you have daemon AND client configured
(respectively with 127.0.0.1:3128 and 172.17.0.1:3128 adresses)

- <https://medium.com/@saniaky/configure-docker-to-use-a-host-proxy-e88bd988c0aa>
- <https://nickjanetakis.com/blog/docker-tip-65-get-your-docker-hosts-ip-address-from-in-a-container>

```bash
#cat > /etc/systemd/system/docker.service.d/http-proxy.conf <<EOF
[Service]
Environment="HTTP_PROXY=http://127.0.0.1:3128"
Environment="HTTPS_PROXY=http://127.0.0.1:3128"
Environment="NO_PROXY=localhost,127.0.0.1,172.17.0.1,172.30.1.1"
EOF
# systemctl daemon-reload
# systemctl restart docker

# in ~/.docker/config.json (172.17.0.1 is the address of host in docker container)
{
 "proxies":
 {
   "default":
   {
     "httpProxy": "http://172.17.0.1:3128",
     "httpsProxy": "http://172.17.0.1:3128",
     "noProxy": "localhost"
   }
 }
}
```

## To debug k8s

See <https://kubernetes.io/fr/docs/reference/kubectl/cheatsheet/> for reference
Useful commands :

- `kubectl get pods --all-namespaces`
- `kubectl describe pod <pod-id>`

## minikube and debug commands

If you have a proxy, use

```bash
unset http_proxy
unset https_proxy
unset HTTP_PROXY
unset HTTPS_PROXY
export HTTP_PROXY=http://10.0.2.2:3128 # 10.0.2.2 is the host adress in virtualbox VM
export HTTPS_PROXY=http://10.0.2.2:3128
export NO_PROXY=$NO_PROXY,192.168.99.100
HTTPS_PROXY=http://10.0.2.2:3128 minikube start --docker-env HTTP_PROXY=http://10.0.2.2:3128 \
                --docker-env HTTPS_PROXY=http://10.0.2.2:3128 \
                --docker-env NO_PROXY=192.168.99.0/24
```

For DNS resolution inside minikube VM, you might have to provide resolv.conf before starting
(see <https://minikube.sigs.k8s.io/docs/tasks/sync/)> :

```bash
mkdir -p ~/.minikube/files/etc
cat /etc/resolv.conf > ~/.minikube/files/etc/resolv.conf
```

You can access you pods binded with service NodePort via the url <http://192.168.99.100:NodePort>/
(for example <http://192.168.99.100:31146>)

## kind and debug commands

If you have a proxy, check

```bash
unset http_proxy
unset https_proxy
unset HTTP_PROXY
unset HTTPS_PROXY
export HTTP_PROXY=http://172.17.0.1:3128  # 172.17.0.1 is the host adress in docker-in-docker(dind)
export HTTPS_PROXY=http://172.17.0.1:3128
export NO_PROXY=172.17.0.0/16
kind create cluster
```

- `kind get clusters`
- `kubectl cluster-info --context kind-kind`
- `kubectl describe pod nginx-example` #debug pod deployement
- `kubectl get pods nginx-example` #get general infos on pod
- `kubectl get svc nginx-example` #get bindings for pod
- `kind delete cluster`

You can access your pod using (see <https://github.com/kubernetes-sigs/kind/issues/99>):

- `docker inspect <containerId> | grep IPAddress` => docker guest ip from host
- `kubectl describe svc <serviceId>` => docker guest port from host

(for example <http://172.17.0.2:31352/)>

## Useful resources

- Kompose, tool to convert docker-compose files to k8s YAML (<https://github.com/kubernetes/kompose>)
- k2tf, tool to convert k8s YAML to Terraform HCL (<https://github.com/sl1pm4t/k2tf>)