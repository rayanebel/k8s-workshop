# Lab02: CREATE YOUR FIRST CONTAINERS

## Objective

Learn how to run, stop and remove containers.

## Instructions

- Run a container in the background named container01 using latest ubuntu image.
  
```bash
docker run -t -d --name container01 ubuntu
```

- Run a container in the background named container02 with latest nginx image and bind the port 80 to the port 8080 on your local machine.

```bash
docker run -t -d --name container02 -p 8080:80 nginx
```

- List the containers to make sure the container is running.

```bash
docker ps
```

- On the container01 open a terminal (bash).

```bash
docker exec -it container01 bash
```

- Stop the containers.

```bash
docker stop container01
docker stop container02
```

- Remove the containers.

```bash
docker rm container01
docker rm container02
```