SHELL := /bin/bash
current_dir = $(shell pwd)

all: build test

test:
	go fmt
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

build:
	go mod download
	go build

clean:
	rm -rf currency-conversion

container:
	docker build . -t currency-conversion:$(shell git rev-parse --short HEAD)
	docker tag currency-conversion:$(shell git rev-parse --short HEAD) currency-conversion:latest
