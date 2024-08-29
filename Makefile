# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@templ generate
	@npx tailwindcss -i public/css/input.css -o public/css/output.css --minify
	@go build -o main.exe cmd/app/main.go

# Run the application
run:
	@go run cmd/app/main.go



# Test the application
test:
	@echo "Testing..."
	@go test ./... -v



# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload

watch:
	@air


.PHONY: all build run test clean watch
