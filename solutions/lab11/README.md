# LAB11: CREATE YOUR FIRST VOLUMES

## Objective

Learn how to create ephemeral and persistent volumes.

## Instructions

- From the previous `podinfo` pod, create an emptyDir volume and mount it in the path `/etc/foo`.

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
        - name: data
          mountPath: /etc/foo
        ports:
        - containerPort: 9898
          name: http
          protocol: TCP
      volumes:
        - name: secrets
          secret:
            secretName: mysecret
        - name: data
          emptyDir: {}
```

- Open a shell on the podinfo pod and add some data in `/etc/foo` _(e.g create a file)_.

```bash
kubectl -n podinfo exec -it <POD_NAME> bash

$ ls /etc/foo
$ echo "test test test" > /etc/foo/data.txt
$ exit
```

- Exit from the pod and restart it.

```bash
kubectl -n podinfo rollout restart deployment podinfo
```

- Open a shell on the podinfo pod and check the content of your volume in `/etc/foo`. Is your data still present?

```bash
kubectl -n podinfo exec -it <POD_NAME> bash

$ ls /etc/foo
```
