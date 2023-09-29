package game

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/async"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/config"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/interaction/applier"
	"github.com/kjkondratuk/goblins-and-gold/model/world"
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

	var w *world.WorldDefinition
	var p actors.Player
	async.InParallel(func() {
		wl := config.Read[world.WorldDefinition]("./data/worlds/test_world.yaml", "./data/monsters.yaml")
		w = &wl
		pterm.Success.Println("World loaded.")
		start.Increment()
	}, func() {
		ps := config.Read[actors.PlayerParams]("./data/test_player.yaml")
		p = actors.NewPlayer(ps)
		pterm.Success.Println("Player loaded.")
		start.Increment()
	})

	s := state.New(ux.NewPromptUiLib(), p, w.StartRoom, w)
	pterm.Success.Println("Game state initialized.")
	start.Increment()

	qc := command.NewQuitCommand(exit)
	lc := command.NewLookCommand()

	cmds := []command.Command{
		command.NewGoCommand(qc, lc),
		command.NewInteractCommand(
			container.NewContainerController(),
			applier.InteractionApplier{},
		),
		command.NewStatsCommand(),
		lc,
		qc,
	}

	if isDebug {
		pterm.EnableDebugMessages()
		cmds = append(cmds, command.NewDebugCommand())
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
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: ", r)
		}

		for {
			// TODO : figure out how to implement this completer
			// prompt the user and read the input
			line := prompt.Input(" >> ", /*func(document prompt.Document) []prompt.Suggest {
					return []prompt.Suggest{}
				}, */command.DefaultCommandCompleter{}.Completer(app), prompt.OptionPrefixTextColor(prompt.Yellow))

			// tokenize the command arguments for processing
			args := strings.Split(line, " ")

			err := app.Run(s, args...)
			if err != nil {
				pterm.Error.Println(err)
			}
		}
	}()
}
