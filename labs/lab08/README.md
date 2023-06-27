# LAB 08: EXPOSE YOUR FIRST APPLICATION

## Objective

Learn how to expose an application.

## Instructions

- From the previous deployment in the namespace `podinfo`
- How do you associate one or more pods with a service ?
- Create a service in `podinfo` namespace that exposes port `9898` and redirects traffic to port `9898` of your pod(s). The service should be of type `LoadBalancer` and should match the pod(s) running the podinfo application.
- List all the services in the namespace podinfo
- Print the details of the service with `kubectl describe service <SERVICE_NAME>`
- Open your browser and try to access to the application on port 9898.