# Steps

See <https://kind.sigs.k8s.io/docs/user/ingress/>

kind ingress with Contour :

- `kubectl apply -f https://projectcontour.io/quickstart/contour.yaml`
- `kubectl patch daemonsets -n projectcontour envoy -p '{"spec":{"template":{"spec":{"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}'`

- `kubectl apply -f ingress-example.yml` # example usage

## To validate

### should output "foo"

curl localhost/foo

### should output "bar"

curl localhost/bar
