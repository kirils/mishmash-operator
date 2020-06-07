# mishmash-operator
Mishmash Operator: http://mishmash.io

Image repository: https://quay.io/repository/kirils/mishmash-operator

## Instalation

Clone with HTTPS
https://github.com/kirils/mishmash-operator.git

Run as a Deployment inside the cluster
```bash
$ kubectl create -f deploy/service_account.yaml
$ kubectl create -f deploy/role.yaml
$ kubectl create -f deploy/role_binding.yaml
$ kubectl create -f deploy/operator.yaml
```
