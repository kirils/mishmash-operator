# mishmash-operator
Mishmash Operator: http://mishmash.io

Image repository: https://quay.io/repository/kirils/mishmash-operator

## Instalation

Clone with HTTPS
https://github.com/kirils/mishmash-operator.git

### Run as a Deployment inside the cluster:
```bash
$ kubectl create -f deploy/service_account.yaml
$ kubectl create -f deploy/role.yaml
$ kubectl create -f deploy/role_binding.yaml
$ kubectl create -f deploy/operator.yaml
```
Verify that the mishmash-operator is up and running:
```bash
$ kubectl get deployment
NAME                     DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
mishmash-operator       1         1         1            1           1m
```
### Run locally outside the cluster
```bash
export OPERATOR_NAME=mishmash-operator
```

```bash
$ operator-sdk run local --watch-namespace=default
2020/06/07 22:15:30 Go Version: go1.10.2
2020/06/07 22:15:30 Go OS/Arch: darwin/amd64
2020/06/07 22:15:30 operator-sdk Version: 0.0.6+git
2020/06/07 22:15:30 Registering Components.
2020/06/07 22:15:30 Starting the Cmd.
```

## Create a Mishmash CR
Create the example Mishmash CR that was generated at deploy/crds/mishmash.io_v1alpha1_mishmash_cr.yaml:

```bash
$ cat deploy/crds/mishmash.io_v1alpha1_mishmash_cr.yaml
apiVersion: "mishmash.io.com/v1alpha1"
kind: "Mishmash"
metadata:
  name: "example-mishmash"
spec:
  size: 3
```

```bash
$ kubectl apply -f deploy/crds/mishmash.io_v1alpha1_mishmash_cr.yaml

Ensure that the mishmash-operator creates the deployment for the CR:

```bash
$ kubectl get deployment
NAME                     DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
mishmash-operator       1         1         1            1           2m
example-mishmash        3         3         3            3           1m
```

Check the pods and CR status to confirm the status is updated with the mishnash pod names:

```bash
$ kubectl get pods
NAME                                  READY     STATUS    RESTARTS   AGE
example-mishmash-6fd7c98d8-7dqdr     1/1       Running   0          1m
example-mishmash-6fd7c98d8-g5k7v     1/1       Running   0          1m
example-mishmash-6fd7c98d8-m7vn7     1/1       Running   0          1m
mishmash-operator-7cc7cfdf86-vvjqk   1/1       Running   0          2m

$ kubectl get mishmash/example-mishmash -o yaml
apiVersion: mishmash.io/v1alpha1
kind: Mishmash
metadata:
  clusterName: ""
  creationTimestamp: 2020-03-31T22:51:08Z
  generation: 0
  name: example-mishmash
  namespace: default
  resourceVersion: "245453"
  selfLink: /apis/mishmash.io/v1alpha1/namespaces/default/mishmashes/example-mishmash
  uid: 0026cc97-3536-11e8-bd83-0800274106a1
spec:
  size: 3
status:
  nodes:
  - example-mishmash-6fd7c98d8-7dqdr
  - example-mishmash-6fd7c98d8-g5k7v
  - example-mishmash-6fd7c98d8-m7vn7
```

### Update the size
Change the spec.size field in the mishmash CR from 3 to 4 and apply the change:

```bash
$ cat deploy/crds/mishmash.io_v1alpha1_mishmash_cr.yaml
apiVersion: "mishmash.io/v1alpha1"
kind: "Mishmash"
metadata:
  name: "example-mishmash"
spec:
  size: 4
```

```bash
$ kubectl apply -f deploy/crds/mishmash.io_v1alpha1_mishmash_cr.yaml
```

Confirm that the operator changes the deployment size:

```bash
$ kubectl get deployment
NAME                 DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
example-mishmash    4         4         4            4           5m
```

### Cleanup
Clean up the resources:

```bash
$ kubectl delete -f deploy/crds/mishmash.io_v1alpha1_mishmash_cr.yaml
$ kubectl delete -f deploy/operator.yaml
$ kubectl delete -f deploy/role_binding.yaml
$ kubectl delete -f deploy/role.yaml
$ kubectl delete -f deploy/service_account.yaml
```
