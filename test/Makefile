clean:  ## Clean test artifacts
	@rm -f coverage.out

.PHONY: test
test: clean  ## Test go modules
	@go test -v ./... -covermode=count -coverprofile=coverage.out
	@go tool cover -func=coverage.out -o=coverage.out
