.PHONY: build test test-cover install-lint lint lint-fix

GOLANGCI_LINT_VERSION := v2.9.0
GOLANGCI_LINT_BIN := ./bin/golangci-lint

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

test:
	go test -v ./...

test-cover:
	go tool gocovreport run --color=always ./...

install-lint:
	curl -sSfL https://golangci-lint.run/install.sh | \
		sh -s -- -b ./bin $(GOLANGCI_LINT_VERSION)

lint:
	$(GOLANGCI_LINT_BIN) run

lint-fix:
	$(GOLANGCI_LINT_BIN) run --fix
