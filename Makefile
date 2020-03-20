.PHONY : all

all: test build lint

build :
	go build -v
test :
	go test -v ./...

lint: 
	golangci-lint run ./...