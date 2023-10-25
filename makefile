.PHONY: default run build test docs clean

# Variables
APP_NAME=finance
APP_VERSION=v1.0.0

# Tasks
default: run-with-docker 

run: 
	@go run main.go

run-with-docker:
	@swag init
	@go run main.go

build:
	@go build -o ${APP_NAME} main.go

test:
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -f ${APP_NAME}
	@rm -rf ./docs
