package game

import (
	"fmt"
	"github.com/chzyer/readline"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/app/commands"
	"github.com/kjkondratuk/goblins-and-gold/app/config"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/player"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
	"os"
	"strings"
	"syscall"
)

const (
	defaultApp = "goblins-and-gold"
)

var (
	defaultAppArr = []string{defaultApp}
)

func Run(appArgs []string, exit chan os.Signal) {
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
			Usage:       "Look at your surroundings",
			Description: "Look at your surroundings",
			Category:    "Info",
			Action:      commands.Look(&s),
		},
		{
			Name:        "go",
			Aliases:     []string{"g"},
			Usage:       "Travel down a path",
			Description: "Travel down a path",
			ArgsUsage:   "[location number]",
			Category:    "Actions",
			Action:      commands.Go(&s, &w),
		}, {
			Name:        "interact",
			Aliases:     []string{"i"},
			Usage:       "Interact with your surroundings",
			Description: "Interact with your surroundings",
			Category:    "Actions",
			Action:      commands.Interact(&s, &w),
		},
		{
			Name:     "stat",
			Aliases:  []string{"s"},
			Usage:    "Interrogate your player stats",
			Category: "Info",
			Action:   commands.Stats(&s),
		},
		{
			Name:    "quit",
			Aliases: []string{"q"},
			Usage:   "Quit the game",
			Action: func(c *cli.Context) error {
				// TODO : should probably write game state here before exiting so we can resume if we want
				fmt.Println("Quitting...")
				exit <- syscall.SIGTERM
				return nil
			},
		},
	}

	// if we're debugging add some additional debug commands that spoil the magic
	if isDebug {
		commands = append(commands, cli.Command{
			Name:     "world",
			Aliases:  []string{"w"},
			Usage:    "Print general info about the world.",
			Category: "Debug",
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
			return cli.ShowAppHelp(c)
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
}
