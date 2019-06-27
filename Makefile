SHELL := /usr/bin/env bash
CWD := $(shell pwd)
BIN := consul-operator

.PHONY: clean

all: $(BIN)

$(BIN): controller.go main.go 
	go build -o $(BIN) main.go controller.go

clean:
	rm -f $(BIN)
