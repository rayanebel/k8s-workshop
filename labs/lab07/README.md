# LAB 07: CREATE YOUR FIRST DEPLOYMENT

## Objective

Learn how to create a deployment, update and scale an application.

## Instructions

- Create a namespace called podinfo
- Create a deployment called podinfo in the `podinfo` namespace with the following settings:
  - Image: `stefanprodan/podinfo:6.3.0`
  - Replicas: 1
  - Labels:
    - `app=podinfo`
  - Ports:
    - `port=9898`, `name=http`, `protocol=TCP`
- List all the deployment in the namespace podinfo.
- List all the pod in the namespace podinfo.
- Update the image tag from 6.3.0 to 6.4.0 and redeploy the application. What do you see?
- Check the update history by using kubectl. Which command will you run?
- Rollback your deployment from 6.4.0 to 6.3.0 by using ONLY kubectl
- Scale your deployment to 5 replicas and check the result.
