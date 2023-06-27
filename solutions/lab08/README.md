# LAB 08: EXPOSE YOUR FIRST APPLICATION

## Objective

Learn how to expose an application.

## Instructions

- From the previous deployment in the namespace `podinfo`
- How do you associate one or more pods with a service ?

By using labels and selectors.

- Create a service in `podinfo` namespace that exposes port `9898` and redirects traffic to port `9898` of your pod(s). The service should be of type `LoadBalancer` and should match the pod(s) running the podinfo application.

```bash
kubectl apply -f - <<EOF
apiVersion: v1
kind: Service
metadata:
  name: podinfo
  namespace: podinfo
spec:
  selector:
    app: podinfo
  ports:
    - name: http
      protocol: TCP
      port: 9898
      targetPort: 9898
EOF
```

- List all the services in the namespace podinfo

```bash
kubectl get services -n podinfo
```

- Print the details of the service with `kubectl describe service <SERVICE_NAME>`

```bash
kubectl describe service podinfo -n podinfo
```

- Open your browser and try to access to the application on port 9898.

```bash
curl http://localhost:9898
```