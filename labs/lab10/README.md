# LAB10: MANAGE YOUR SENSITIVE DATA WITH SECRETS

## Objective

Learn how to use secrets to manage your sensitive data.

## Instructions

- Create a secret called mysecret with values `'password=mysecretpassword'` in the namespace `podinfo`.
- Display the content of the secret with kubectl ? What can you see ?
- Load this secret as env variables into the podinfo pod.
- Open a shell in the podinfo pod and check the environment variables. Can you see the variables defined in the secrets ?
- `BONUS`: Load the same secret as a volume inside the podinfo pod.
