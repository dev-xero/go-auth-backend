format:
	go fmt ./...

tidy:
	go mod tidy

server: format tidy
	go run server.go

.PHONY: format tidy