# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.DEFAULT_GOAL := help

.PHONY: help
# Put it first so that "make" without argument is like "make help".
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-32s-\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: format
format:  ## Format go modules
	@go fmt .

clean:  ## Clean test artifacts
	@rm -f coverage.out

.PHONY: test
test: clean  ## Test go modules
	@go test -v ./... -covermode=count -coverprofile=coverage.out
	@go tool cover -func=coverage.out -o=coverage.out
