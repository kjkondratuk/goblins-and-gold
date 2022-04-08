package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/config"
	"github.com/kjkondratuk/goblins-and-gold/player"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/world"
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

	isDebug := strings.EqualFold(os.Getenv("DEBUG"), "true")

	fmt.Println("Loading world...")
	w := config.Read[world.World]("./data/test_world.yaml")

	fmt.Println("Loading player...")
	p := config.Read[player.Player]("./data/test_player.yaml")
	sr, _ := w.Room(w.StartRoom)
	s := state.GameState{
		Player:   p,
		CurrRoom: sr,
	}

	fmt.Println("Configuring application...")
	commands := []cli.Command{
		{
			Name:        "look",
			Aliases:     []string{"l"},
			Description: "Look at your surroundings",
			Action: func(c *cli.Context) error {
				//fmt.Println("Looking around...")
				fmt.Println(s.CurrRoom.Description)
				return nil
			},
		},
		{
			Name:        "stat",
			Aliases:     []string{"s"},
			Description: "Interrogate your player stats",
			Action: func(c *cli.Context) error {
				ps, _ := yaml.Marshal(s.Player)
				fmt.Println(pterm.Green(string(ps)))
				return nil
			},
		},
		{
			Name:        "quit",
			Aliases:     []string{"q"},
			Description: "Quit the game",
			Action: func(c *cli.Context) error {
				fmt.Println("Quitting...")
				exit <- syscall.SIGTERM
				return nil
			},
		},
	}

	// if we're debugging add some additional debug commands that spoil the magic
	if isDebug {
		commands = append(commands, cli.Command{
			Name:        "world",
			Aliases:     []string{"w"},
			Description: "Print general info about the world.",
			Action: func(c *cli.Context) error {
				ws, _ := yaml.Marshal(w)
				fmt.Println(pterm.Green(string(ws)))
				return nil
			},
		})
	}

	app := &cli.App{
		Name: defaultApp,
		Action: func(c *cli.Context) error {
			fmt.Println("Running app...")
			return nil
		},
		Commands: commands,
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

	fmt.Println("Starting game...")

	// REPL Loop
	go func() {
		for {
			// prompt the user and read the input
			line, err := rl.Readline()
			if err != nil {
				fmt.Printf("Error: %+v\n", err)
				exit <- syscall.SIGTERM
			}

			// tokenize the command arguments for processing
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
