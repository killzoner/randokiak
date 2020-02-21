# How to use local kubernetes

You will find useful links about various parts of the stack to help you start up a local k8s (Kubernetes) cluster and manager it.

This config relies on `kind`to leverage existing docker installation (<https://github.com/kubernetes-sigs/kind)>, but you can replace minikube with `minikube` if you prefer so.

Documentation on both of them:

- <https://kubernetes.io/fr/docs/setup/learning-environment/minikube/>

Documentation about Terraform:

- <https://www.terraform.io/docs/providers/kubernetes/guides/getting-started.html>
- <https://learn.hashicorp.com/terraform/getting-started/intro>
- <https://itmecho.com/managing-kubernetes-with-terraform/>
- <https://www.terraform.io/docs/configuration/>

## Tips about local Docker config and accesses to the Docker daemon

With Docker and a proxy conf, make sure you have daemon AND client configured
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

For ingress support, use:

```bash
cat <<EOF | kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
        authorization-mode: "AlwaysAllow"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF
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

## Use local images in Minikube (no special config for kind)

Follow these steps:

- Set the environment variables with eval $(minikube docker-env)
- Build the image with the Docker daemon of Minikube (eg docker build -t my-image .)
- Set the image in the pod spec like the build tag (eg my-image)
- Set the imagePullPolicy to Never, otherwise Kubernetes will try to download the image.

## Use local images in kind

See <https://kind.sigs.k8s.io/docs/user/quick-start/#loading-an-image-into-your-cluster>

## Useful resources

- Kompose, tool to convert docker-compose files to k8s YAML (<https://github.com/kubernetes/kompose>)
- k2tf, tool to convert k8s YAML to Terraform HCL (<https://github.com/sl1pm4t/k2tf>)
