SHELL = /bin/sh

# all: Run tests and perform build
.PHONY: all
all: deps test
	go build

# help: Print help information
.PHONY: help
help:
	@echo ">> Help info for supported targets:"
	@grep -E '^# [-a-z./]+:' Makefile | grep -v https:// | sed -e 's|#|   make|g' | sort

# coverage: Display code coverage in html
.PHONY: coverage
coverage: test
	@echo ">> Rendering code coverage"
	go tool cover -html=coverage.txt
	@echo echo "Success 👍"

# generate: Run go generate for all packages
.PHONY: generate
generate:
	@echo ">> Running codegen"
	go generate -v

# test: Run go test
.PHONY: test
test: generate
	@echo ">> Running tests"
	go test -race -coverprofile=coverage.txt -covermode=atomic -v
	@echo echo "Success 👍"

# deps: Install dependencies
.PHONY: deps
deps:
	@echo ">> Installing dependencies"
	@go get github.com/GeertJohan/go.rice/rice
	@go get -t -v ./...
	@echo echo "Success 👍"

# vet: Run go vet
.PHONY: vet
vet: generate
	@echo ">> Running go vet"
	go vet -x
	@echo echo "Success 👍"
