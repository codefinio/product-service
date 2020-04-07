
.PHONY: build
build:
	go build -o bin/server cmd/**/*.go

.PHONY: start
start:
	go run cmd/server/server.go