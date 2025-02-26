swag:
	swag init -d ./internal -g ../cmd/main.go -o ./cmd/docs
.PHONY: swag

build:
	go build -o ./cmd/main ./cmd/main.go
.PHONY: build

run:
	go run ./cmd/main.go
.PHONY: run
