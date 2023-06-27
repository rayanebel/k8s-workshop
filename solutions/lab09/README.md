# LAB 09: MANAGE YOUR DATA WITH CONFIGMAPS

## Objective

Learn how to use secrets to manage your sensitive data.

## Instructions

- Create a configmap called `configs` with values `'var1=val1'`, `'var2=val2'` in the namespace `podinfo`.
- List all the configmaps in the namespace `podinfo`.

```bash
kubectl -n podinfo get configmaps
```

- Display the details of teh configmap with `kubectl`

```bash
kubectl -n podinfo describe configmap configs
```
kubectl -n podinfo exec -it <POD_NAME> bash

```bash
kubectl -n podinfo create cm configs --from-literal=var1=val1 --from-literal=var2=val2
```

- Load this configmap as env variables into the podinfo pod.

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
        envFrom:
        - configMapRef:
            name: configs
        ports:
        - containerPort: 9898
          name: http
          protocol: TCP
EOF
```

- Open a shell in the `podinfo` pod and check the environment variables. Can you see the variables defined in the configmap ?

```bash
kubectl -n podinfo exec -it <POD_NAME> sh

$> env
```

- `BONUS`: Load the same configmap as a volume inside the podinfo pod.

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
        volumeMounts:
        - name: config
          mountPath: /etc/config
        ports:
        - containerPort: 9898
          name: http
          protocol: TCP
      volumes:
      - name: config
        configMap:
          name: configs
EOF
```
