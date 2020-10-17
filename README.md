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
```

- Open the results in the browser:

```shell
$ go tool cover -html=coverage.out
<Opens Browser>
```
