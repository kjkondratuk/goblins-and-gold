GO=go
MAKE=make

GENERATED_XML_DIR=$(shell pwd)/internal/generated/xml/game

generate:
	xgen -i worlds/schema.xml -o $(GENERATED_XML_DIR) -l Go

build:
	go build -o $(shell pwd)/bin/seed $(shell pwd)/cmd/world/seed.go
	go build -o $(shell pwd)/bin/gng $(shell pwd)/cmd/game/gng.go