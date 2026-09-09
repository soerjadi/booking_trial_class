.PHONY: run tidy build

run:
	go run ./cmd/rest

build:
	go build -o bin/booking ./cmd/rest

tidy:
	go mod tidy
