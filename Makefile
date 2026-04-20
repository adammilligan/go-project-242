.PHONY: build test install-tools test-cover lint lint-fix

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

test:
	go test -v ./...

install-tools:
	go install github.com/adammilligan/gocovreport/cmd/gocovreport@v0.0.0-20260420084811-3ad22bb03482

test-cover: install-tools
	gocovreport run --color=always ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix