#include .env
# https://ieftimov.com/post/golang-package-multiple-binaries/
# https://github.com/gobuffalo/packr/blob/master/packr/main.go
# checkout godag?
GOROOT = .
SIMCO_BIN = "./bin"
SIMCO_CMD = "./cmd"


build:
	go build -v -o $(SIMCO_BIN) $(SIMCO_CMD)/util_status_resource

run:
	./bin/util_status_resource

test:
	go test
