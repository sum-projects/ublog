# syntax=docker/dockerfile:1

FROM golang:1.18.5-alpine3.16

WORKDIR /app

RUN apk add --update \
    curl \
    && rm -rf /var/cache/apk/*

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o  blog-api cmd/blog-api/main.go

EXPOSE 8080

CMD [ "air" ]