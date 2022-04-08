package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	defaultApp = "goblins-and-gold"
)

var (
	defaultAppArr = []string{defaultApp}
)

func main() {
	// setup exit listener
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// TODO : create a collections of "applications"
	app := &cli.App{
		Name: defaultApp,
		Action: func(c *cli.Context) error {
			fmt.Println("Running app...")

			return nil
		},
		Commands: []cli.Command{
			{
				Name:    "look",
				Aliases: []string{"l"},
				Action: func(c *cli.Context) error {
					fmt.Println("Looking around...")
					return nil
				},
			},
			{
				Name:    "quit",
				Aliases: []string{"q"},
				Action: func(c *cli.Context) error {
					fmt.Println("Quitting...")
					exit <- syscall.SIGTERM
					return nil
				},
			},
		},
	}

	ac := readline.NewPrefixCompleter()

	rl, err := readline.NewEx(&readline.Config{
		Prompt:       pterm.LightYellow(" >> "),
		AutoComplete: ac,
	})
	if err != nil {
		fmt.Println("Error starting repl.  Exiting.")
		exit <- syscall.SIGTERM
	}

	go func() {
		for {
			line, err := rl.Readline()
			if err != nil {
				fmt.Printf("Error: %+v\n", err)
				exit <- syscall.SIGTERM
			}
			args := strings.Split(line, " ")

			// prepend the "app" qualifier to all calls to execute the default application
			args = append(defaultAppArr, args...)

			err = app.Run(args)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	sig := <-exit
	fmt.Printf("%s received, exiting...\n", sig)
}
