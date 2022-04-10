package main

import (
	"github.com/kjkondratuk/goblins-and-gold/app/game"
	"github.com/pterm/pterm"
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
	pterm.Info.Printf("%s received, exiting...\n", sig)
}
