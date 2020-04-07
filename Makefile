
.PHONY: build
build:
	go build -o bin/server cmd/server/*.go
	go build -o bin/migration cmd/migration/*.go

.PHONY: start
start:
	go run cmd/server/server.go

.PHONY: test
test:
	go test -v ./... -cover