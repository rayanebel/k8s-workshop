# LAB 06: CREATE YOUR FIRST POD

## Objective

Learn how to create a namespace and a pod.

## Instructions

- Create a namespace called `podinfo`

```bash
kubectl create namespace podinfo
```

- List all the available namespaces.

```bash
kubectl get namespaces
```

- Create a pod named `podinfo` in the namespace `podinfo` with the following settings:
  - Image: `stefanprodan/podinfo:6.3.0`
  - Ports:
    - `port=9898`, `name=http`, `protocol=TCP`

```bash
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: podinfo
  namespace: podinfo
spec:
  containers:
  - name: podinfo
    image: stefanprodan/podinfo:6.3.0
    ports:
    - containerPort: 9898
      name: http
      protocol: TCP
EOF
```

- Verify if your pod is running. Which command will you run?

```bash
kubectl get pods -n podinfo
```

- Displays detailed status of your pod with `kubectl`. Which command will you run?

```bash
kubectl describe pods podinfo -n podinfo
```

- Open a shell inside your pod. Which command will you run?

```bash
kubectl exec -it podinfo -n podinfo -- sh
```

- Try to edit the port number from `9898` to `9899`. What do you see?

```bash
kubectl apply -f - <<EOF
apiVersion: v1
kind: Pod
metadata:
  name: podinfo
  namespace: podinfo
spec:
  containers:
  - name: podinfo
    image: stefanprodan/podinfo:6.3.0
    ports:
    - containerPort: 9899
      name: http
      protocol: TCP
EOF
```

We can't perform an upgrade on a pod directly. We have to destroy and recreate the pod. Automated upgrade are a feature available in Deployments.

- Delete the `podinfo` pod 

```bashkubectl -n podinfo rollout history deployment podinfo
kubectl -n podinfo delete pod podinfo
```

- Delete the `podinfo` namespace

```bash
kubectl delete ns podinfo
```