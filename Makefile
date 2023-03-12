GO=go
PROJ_DIR=$(shell pwd)
BIN=.$(PROJ_DIR)/bin/goblins-and-gold

.PHONY: run
run:
	$(GO) build -o $(BIN) .
	$(BIN)

.PHONY: generate
generate:
	$(GO) generate ./...
