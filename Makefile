CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest
DOCKERFILE=Dockerfile

run:
	go run cmd/main.go

linter:
	golangci-lint run

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go
