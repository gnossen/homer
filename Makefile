all: bin/go-dm

SRC=$(shell find src -name '*.go')

bin/go-dm:  $(SRC)
	mkdir -p bin pkg
	go install github.com/gnossen/homer
	go test github.com/gnossen/homer

.PHONY: format
format:
	go fmt github.com/gnossen/...

.PHONY: clean
clean:
	rm -rf bin pkg
