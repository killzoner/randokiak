# Steps

- `kubectl apply -f dashboard.yaml`
- `kubectl apply -f dashboard-adminuser.yaml`
- `kubectl apply -f dashboard-cluster-role-binding.yaml`
- `kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')`
- `kubectl proxy`
- go to <http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy>
