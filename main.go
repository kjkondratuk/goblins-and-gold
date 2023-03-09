package main

import (
	"github.com/kjkondratuk/goblins-and-gold/game"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// setup exit listener
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	game.Run(os.Args, exit)

	<-exit
}
