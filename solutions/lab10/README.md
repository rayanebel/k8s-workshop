# LAB10: MANAGE YOUR SENSITIVE DATA WITH SECRETS

## Objective

Learn how to use secrets to manage your sensitive data.

## Instructions

- Create a secret called mysecret with values `'password=mysecretpassword'` in the namespace `podinfo`.

```bash
kubectl -n podinfo create secret generic mysecret --from-literal=password=mysecretpassword
```

- Display the content of the secret with kubectl ? What can you see ?

```bash
kubectl -n podinfo get secrets -o yaml mysecret
```

The secret is encoded in base64. To decode it you can use the following command:

```bash
echo "<SECRET_ENCODED>" | base64 -d
```

- Load this secret as env variables into the podinfo pod.

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
        - secretRef:
            name: mysecret
        ports:
        - containerPort: 9898
          name: http
          protocol: TCP
EOF 
```

- Open a shell in the podinfo pod and check the environment variables. Can you see the variables defined in the secrets ?

```bash
kubectl -n podinfo exec -it <POD_NAME> bash

$> env
```

- `BONUS`: Load the same secret as a volume inside the podinfo pod.

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
        - secretRef:
            name: mysecret
        volumeMounts:
        - name: secrets
          mountPath: /etc/secrets
          readOnly: true
        ports:
        - containerPort: 9898
          name: http
          protocol: TCP
      volumes:
        - name: secrets
          secret:
            secretName: mysecret
```
