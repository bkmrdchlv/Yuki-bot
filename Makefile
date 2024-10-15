run:
	go run cmd/main.go

linter:
	golangci-lint run

build:
	go build cmd/main.go
