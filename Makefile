.PHONY: all build run test clean lint fmt wire

all: fmt lint wire build test

build: wire
	go build -o bin/goscaff ./cmd/goscaff

run: build
	./bin/goscaff -c configs/config.dev.yaml

test:
	go test ./...

clean:
	rm -rf bin/

lint:
	golangci-lint run ./...

fmt:
	go fmt ./...

wire:
	wire gen ./...
