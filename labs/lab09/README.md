# LAB 09: MANAGE YOUR DATA WITH CONFIGMAPS

## Objective

Learn how to use secrets to manage your sensitive data.

## Instructions

- Create a configmap called configs with values `'var1=val1'`, `'var2=val2'` in the namespace `podinfo`.
- Load this configmap as env variables into the podinfo pod.
- Open a shell in the `podinfo` pod and check the environment variables. Can you see the variables defined in the configmap ?
- `BONUS`: Load the same configmap as a volume inside the podinfo pod.
