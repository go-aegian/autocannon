.PHONY: build

build: test clean
	go build -o ./build/autocannon ./autocannon.go

test:
	go test ./...

clean:
	go clean ./...
