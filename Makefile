.PHONY: build test install-tools test-cover install-lint lint lint-fix

GOLANGCI_LINT_VERSION := v2.9.0
GOLANGCI_LINT_BIN := ./bin/golangci-lint

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

test:
	go test -v ./...

install-tools:
	go install github.com/adammilligan/gocovreport/cmd/gocovreport@v0.0.0-20260420084811-3ad22bb03482

test-cover: install-tools
	gocovreport run --color=always ./...

install-lint:
	curl -sSfL https://golangci-lint.run/install.sh | \
		sh -s -- -b ./bin $(GOLANGCI_LINT_VERSION)

lint: install-lint
	$(GOLANGCI_LINT_BIN) run

lint-fix:
	$(GOLANGCI_LINT_BIN) run --fix
