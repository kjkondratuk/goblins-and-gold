package game

import (
	"github.com/c-bata/go-prompt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/async"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/config"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
	"os"
	"strings"
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

	var w *state.WorldDefinition
	var p actors.Player
	async.InParallel(func() {
		wl := config.Read[state.WorldDefinition]("./data/worlds/test_world.yaml", "./data/monsters.yaml")
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
	s := state.New(ux.NewPromptUiLib(), p, &sr, w)
	pterm.Success.Println("Game state initialized.")
	start.Increment()

	//cmds := []cli.Command{
	//	look.New(s),
	//	_go.New(s),
	//	interact.New(s),
	//	stats.New(s),
	//	{
	//		Name:        "quit",
	//		Aliases:     []string{"q"},
	//		Usage:       "Quit the game",
	//		Description: "Quit the game",
	//		Action: func(c *cli.Context) error {
	//			// TODO : should probably write game state here before exiting so we can resume if we want
	//			pterm.Info.Println("Quitting...")
	//			exit <- syscall.SIGTERM
	//			return nil
	//		},
	//	},
	//}

	// if we're debugging add some additional debug cmds that spoil the magic
	//if isDebug {
	//	pterm.EnableDebugMessages()
	//	cmds = append(cmds, cli.Command{
	//		Name:        "world",
	//		Aliases:     []string{"w"},
	//		Usage:       "Print general info about the world.",
	//		Description: "Print general info about the world.",
	//		Category:    "Debug",
	//		Action: func(world *state.WorldDefinition) cli.ActionFunc {
	//			return func(c *cli.Context) error {
	//				ws, _ := yaml.Marshal(world)
	//				pterm.Debug.Println(pterm.Green(string(ws)))
	//				return nil
	//			}
	//		}(w),
	//	})
	//}

	//app := &cli.App{
	//	Name: defaultApp,
	//	Action: func(c *cli.Context) error {
	//		return cli.ShowAppHelp(c)
	//	},
	//	Commands: cmds,
	//}

	qc := command.NewQuitCommand(exit)

	cmds := []command.Command{
		command.NewGoCommand(qc),
		command.NewStatsCommand(),
		command.NewLookCommand(),
		qc,
	}

	if isDebug {
		pterm.EnableDebugMessages()
		cmds = append(cmds, command.NewDebugCommand(command.NewWorldCommand()))
	}

	app := command.NewApp(
		"goblins-and-gold",
		"A daring game of text-based dungeon crawling",
		cmds...,
	)

	pterm.Success.Println("Application configured.")
	start.Increment()

	pterm.Info.Println("Starting game...")

	// REPL Loop
	go func() {
		// TODO : do we need to recover here?

		for {
			// prompt the user and read the input
			line := prompt.Input(" >> ", func(document prompt.Document) []prompt.Suggest {
				return []prompt.Suggest{}
			}, /*completer(app.Commands())*/ prompt.OptionPrefixTextColor(prompt.Yellow))

			// tokenize the command arguments for processing
			args := strings.Split(line, " ")

			err := app.Run(s, args...)
			if err != nil {
				pterm.Error.Println(err)
			}
		}
	}()
}

//func completer(cList []command.Command) prompt.Completer {
//	return func(d prompt.Document) []prompt.Suggest {
//		var s []prompt.Suggest
//		for _, c := range cList {
//			s = append(s, prompt.Suggest{Text: c.Name(), Description: c.Description()})
//		}
//
//		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
//	}
//}
