run:
	@go run cmd/api/main.go

test:
	@go test ./...

test-cov:
	@go test ./... -covermode=atomic