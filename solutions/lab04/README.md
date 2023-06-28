# Lab04: CREATE YOUR OWN CONTAINER IMAGE

## Objective

Learn how to create your own container image from a real application developed in Golang.

## Instructions

- Let’s build a small application _(live coding if time)_.

The source code is available [here](./application/)

- Create the Dockerfile for the application.

```dockerfile
FROM golang:1.20 as builder

WORKDIR /app
COPY *.go .
COPY go.* .

RUN CGO_ENABLED=0 GOOS=linux go build -o blog-api

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/blog-api ./

EXPOSE 3000

CMD ["./blog-api"]
```

You can retrieve the Dockerfile [here](./application/Dockerfile)

- Create the container image from the Dockerfile.

```bash
docker build -t blog-api:1.0.0 -f application/Dockerfile .
```

- Create a container with the image built in the previous step.

```bash
docker run -t -d --name blog-api -p 3000:3000 blog-api:1.0.0
```

- Update the codebase and build a newest version of the image.

After updating some line in the codebase you can build a new version with the following command

```bash
docker build -t blog-api:2.0.0 -f application/Dockerfile .
```

- Recreate the container with the newest image.
```bash
docker stop blog-api && docker rm blog-api
docker run -t -d --name blog-api -p 3000:3000 blog-api:2.0.0 
```
