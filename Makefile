GO=go
MAKE=make
PWD=$(shell pwd)

build:
	go build -o ${PWD}/bin/seed ${PWD}/cmd/world/seed.go
	go build -o ${PWD}/bin/gng ${PWD}/cmd/game/gng.go