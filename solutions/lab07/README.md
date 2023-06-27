# LAB 07: CREATE YOUR FIRST DEPLOYMENT

## Objective

Learn how to create a deployment, update and scale an application.

## Instructions

- Create a namespace called podinfo

```bash
kubectl create ns podinfo
```

- Create a deployment called podinfo in the `podinfo` namespace with the following settings:
  - Image: `stefanprodan/podinfo:6.3.0`
  - Replicas: 1
  - Labels:
    - `app=podinfo`
  - Ports:
    - `port=9898`, `name=http`, `protocol=TCP`

```bash
kubectl apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  name: podinfo
  namespace: podinfo
  labels:
    app: podinfo
spec:
  replicas: 1
  selector:
    matchLabels:
      app: podinfo
  template:
    metadata:
      labels:
        app: podinfo
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

- List all the deployment in the namespace podinfo.

```bash
kubectl -n podinfo get deploy
```

- List all the pod in the namespace podinfo.

```bash
kubectl -n podinfo get pods
```

- Update the image tag from 6.3.0 to 6.4.0 and redeploy the application. What do you see?

```bash
kubectl -n podinfo set image deploy podinfo podinfo=stefanprodan/podinfo:6.4.0
```

- Check the update history by using kubectl. Which command will you run?

```bash
kubectl -n podinfo rollout history deployment podinfo
```

- Rollback your deployment from 6.4.0 to 6.3.0 by using ONLY kubectl

```bash
kubectl -n podinfo rollout undo deployment podinfo
```

- Scale your deployment to 5 replicas and check the result.

```bash
kubectl -n podinfo scale deploy podinfo --replicas=5
```