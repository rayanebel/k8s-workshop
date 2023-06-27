# LAB 06: CREATE YOUR FIRST POD

## Objective

Learn how to create a namespace and a pod.

## Instructions

- Create a namespace called `podinfo`
- List all the available namespaces.
- Create a pod named `podinfo` in the `podinfo` namespace with the following settings:
  - Image: `stefanprodan/podinfo:6.3.0`
  - Labels:
    - `app=podinfo`
  - Ports:
    - `port=9898`, `name=http`, `protocol=TCP`
- Verify if your pod is running. Which command will you run?
- Displays detailed status of your pod with `kubectl`. Which command will you run?
- Open a shell inside your pod. Which command will you run?
- Update the image tag from `6.3.0` to `6.4.0` and redeploy the application. What do you see?
- Delete the `podinfo` pod 
- Delete the `podinfo` namespace
