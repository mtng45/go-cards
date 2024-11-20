.PHONY: build clean run test

# バイナリ名を設定
BINARY_NAME=cards

build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) cmd/main.go

clean:
	@echo "Cleaning..."
	@rm -rf bin/*
	@rm -rf out/*

run: build
	@echo "Running..."
	@./bin/$(BINARY_NAME)

test:
	@echo "Testing..."
	@go test ./...