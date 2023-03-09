GO=go
PROJ_DIR=$(shell pwd)
BIN=.$(PROJ_DIR)/bin/goblins-and-gold


run:
	$(GO) build -o $(BIN) .
	$(BIN)