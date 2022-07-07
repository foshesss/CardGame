all: build test

build:
	go build -o CardGame main.go

run:
	make build
	./CardGame

test:
	go test -v ./cards

clean:
	rm ./CardGame
	rm -rf ./dataStore/*

.PHONY: all build clean
