.PHONY: all build run test clean lint fmt wire

all: fmt lint wire build test

build:
	go build -o bin/goscaff ./cmd/goscaff

run:
	go run cmd/goscaff/main.go

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
