package game

import (
	"github.com/c-bata/go-prompt"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/async"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/command/go"
	"github.com/kjkondratuk/goblins-and-gold/app/config"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
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

	start, _ := pterm.DefaultProgressbar.WithTotal(4).WithTitle("Starting...").Start()

	var w *world.Definition
	var p actors.Player
	async.InParallel(func() {
		wl := config.Read[world.Definition]("./data/worlds/test_world.yaml", "./data/monsters.yaml")
		w = &wl
		pterm.Success.Println("World loaded.")
		start.Increment()
	}, func() {
		ps := config.Read[actors.PlayerParams]("./data/test_player.yaml")
		p = actors.NewPlayer(ps)
		pterm.Success.Println("Player loaded.")
		start.Increment()
	})

	sr, _ := w.Room(w.StartRoom)
	s := &state.State{
		Player:        p,
		CurrRoom:      &sr,
		World:         w,
		SelectBuilder: ux.New(),
	}
	pterm.Success.Println("Game state initialized.")
	start.Increment()

	cmds := []cli.Command{
		{
			Name:        "look",
			Aliases:     []string{"l"},
			Usage:       "Look at your surroundings",
			Description: "Look at your surroundings",
			Category:    "Info",
			Action:      command.Look(s),
		},
		_go.NewGoCommand(s), {
			Name:        "interact",
			Aliases:     []string{"i"},
			Usage:       "Interact with your surroundings",
			Description: "Interact with your surroundings",
			Category:    "Actions",
			Action:      command.Interact(s),
		},
		{
			Name:        "stats",
			Aliases:     []string{"s"},
			Usage:       "Interrogate your player stats",
			Description: "Interrogate your player stats",
			Category:    "Info",
			Action:      command.Stats(s),
		},
		{
			Name:        "quit",
			Aliases:     []string{"q"},
			Usage:       "Quit the game",
			Description: "Quit the game",
			Action: func(c *cli.Context) error {
				// TODO : should probably write game state here before exiting so we can resume if we want
				pterm.Info.Println("Quitting...")
				exit <- syscall.SIGTERM
				return nil
			},
		},
	}

	// if we're debugging add some additional debug cmds that spoil the magic
	if isDebug {
		pterm.EnableDebugMessages()
		cmds = append(cmds, cli.Command{
			Name:        "world",
			Aliases:     []string{"w"},
			Usage:       "Print general info about the world.",
			Description: "Print general info about the world.",
			Category:    "Debug",
			Action: func(world *world.Definition) cli.ActionFunc {
				return func(c *cli.Context) error {
					ws, _ := yaml.Marshal(world)
					pterm.Debug.Println(pterm.Green(string(ws)))
					return nil
				}
			}(w),
		})
	}

	app := &cli.App{
		Name: defaultApp,
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
		Commands: cmds,
	}

	pterm.Success.Println("Application configured.")
	start.Increment()

	pterm.Info.Println("Starting game...")

	// REPL Loop
	go func() {
		for {
			// prompt the user and read the input
			line := prompt.Input(" >> ", completer(app.Commands), prompt.OptionPrefixTextColor(prompt.Yellow))

			// tokenize the command arguments for processing
			args := strings.Split(line, " ")

			// prepend the "app" qualifier to all calls to execute the default application
			args = append(defaultAppArr, args...)

			err := app.Run(args)
			if err != nil {
				pterm.Error.Println(err)
			}
		}
	}()
}

func completer(cList []cli.Command) func(prompt.Document) []prompt.Suggest {
	return func(d prompt.Document) []prompt.Suggest {
		//s := []prompt.Suggest{
		//{Text: "look", Description: "Look around"},
		//{Text: "go", Description: "Travel"},
		//{Text: "interact", Description: "Interact with your surroundings"},
		//{Text: "stats", Description: "View your player stats"},
		//{Text: "quit", Description: "Exit the game"},
		//}
		var s []prompt.Suggest
		for _, c := range cList {
			s = append(s, prompt.Suggest{Text: c.Name, Description: c.Description})
		}

		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}
}
