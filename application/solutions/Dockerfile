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