.PHONY: default run build test docs clean run-with-docs

# Variable
APP_NAME=gopportunities

# Task
default: run-with-docs 

run: 
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs