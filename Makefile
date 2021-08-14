# Build And Development
test:
	@go test ./... -race -coverprofile=coverage.txt -covermode=atomic