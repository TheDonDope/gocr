# gocr

An OCR CLI tool implemented in golang.

## Building

- Build the main command:

```shell
$ go build ./cmd/gocr
<Empty output on build success>
```

## Running Tests

- Run the testsuite with coverage enabled:

```shell
$ go test -coverpkg=all ./... -coverprofile=coverage.out
?       github.com/TheDonDope/gocr/cmd/gocr     [no test files]
?       github.com/TheDonDope/gocr/pkg/config   [no test files]
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```

## Running with Docker

- Build the image:

```shell
$ docker-compose -p app build
db uses an image, skipping
Building app
Step 1/11 : FROM golang:1.15-buster AS builder
 ---> 4a581cd6feb1
Step 2/11 : WORKDIR /src
 ---> Using cache
 ---> 7c57d3937172
Step 3/11 : COPY go.mod go.sum ./
 ---> Using cache
 ---> 53fc8b956733
Step 4/11 : RUN go mod download -x
 ---> Using cache
 ---> e96e6b42503d
Step 5/11 : COPY . ./
 ---> Using cache
 ---> 984b656434d9
Step 6/11 : RUN go build -v -o /bin/gocr cmd/gocr/*.go
 ---> Using cache
 ---> b38ec08f321e

Step 7/11 : FROM debian:buster-slim
 ---> f49666103347
Step 8/11 : RUN set -x && apt-get update &&   DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates &&   rm -rf /var/lib/apt/lists/*
 ---> Using cache
 ---> 75f328a5c1ae
Step 9/11 : WORKDIR /app
 ---> Using cache
 ---> 3540cb5d4a0a
Step 10/11 : COPY --from=builder /bin/gocr ./
 ---> Using cache
 ---> 5d3446c67030
Step 11/11 : CMD ["./gocr"]
 ---> Using cache
 ---> 4a6c7788a63b

Successfully built 4a6c7788a63b
Successfully tagged app_app:latest
```

- Run the app:

```shell
$ docker-compose -p app up -d
Creating app_db_1 ... done
Creating app_app_1 ... done
```
