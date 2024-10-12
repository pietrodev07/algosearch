BINARY=algosearch

build:
	go build -o bin/$(BINARY) cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./searching
