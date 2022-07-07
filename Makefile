all: build

build:
	go build -o CardGame cards/main.go cards/deck.go

run:
	make build
	./CardGame

clean:
	rm ./CardGame

.PHONY: all build clean
