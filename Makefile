all: bin/go-dm

SRC=$(shell find src -name '*.go')

bin/go-dm:  $(SRC)
	mkdir -p bin pkg
	go install github.com/gnossen/go-dm
	go test github.com/gnossen/go-dm

.PHONY: clean
clean:
	rm -rf bin pkg
