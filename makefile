format:
	go fmt ./...

tidy:
	go mod tidy

server: format tidy
	ENVIRONMENT=development go run server.go

.PHONY: format tidy