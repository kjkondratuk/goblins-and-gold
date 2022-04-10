package main

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/app/game"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// setup exit listener
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	game.Run(os.Args, exit)

	sig := <-exit
	fmt.Printf("%s received, exiting...\n", sig)
}
