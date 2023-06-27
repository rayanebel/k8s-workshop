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
kubectl exec -it podinfo -n podinfo -- bash
```

- Update the image tag from `6.3.0` to `6.4.0` and redeploy the application. What do you see?

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
    image: stefanprodan/podinfo:6.4.0
    ports:
    - containerPort: 9898
        name: http
        protocol: TCP
EOF
```

We can't perform an upgrade on a pod directly. We have to destroy and recreate the pod. Automated upgrade are a feature available in Deployments.

- Delete the `podinfo` pod 

```bash
kubectl delete pod podinfo
```

- Delete the `podinfo` namespace

```bash
kubectl delete ns podinfo
```