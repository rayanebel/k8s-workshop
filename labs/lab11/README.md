# LAB11: CREATE YOUR FIRST VOLUMES

## Objective

Learn how to create ephemeral and persistent volumes.

## Instructions

- From the previous `podinfo` pod, create an emptyDir volume and mount it in the path `/etc/foo`.
- Open a shell on the podinfo pod and add some data in `/etc/foo` _(e.g create a file)_.
- Exit from the pod and restart it.
- Open a shell on the podinfo pod and check the content of your volume in `/etc/foo`. Is your data still present?
