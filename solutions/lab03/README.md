# Lab03: WORKING WITH CONTAINER IMAGES

## Objective

Learn how to work with container images.

## Instructions

- List the containers images in your environment.

```bash
docker image ls
```

- Pull the latest `stefanprodan/podinfo` image.
  
```bash
docker pull stefanprodan/podinfo
```

- Run a container in the background named `podinfo` with the image `stefanprodan/podinfo` and bind the port `9898` to the port `9898` on your local machine.

```bash
docker run -t -d --name podinfo -p 9898:9898 stefanprodan/podinfo
```

- Remove the image. Did it work ?
- Do whatever is needed in order to remove the image.

```bash
docker stop podinfo
docker rm podinfo
docker rmi stefanprodan/podinfo
```
