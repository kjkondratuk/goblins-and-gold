package game

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/async"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/config"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/interaction/applier"
	"github.com/kjkondratuk/goblins-and-gold/model/world"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ui"
	"github.com/pterm/pterm"
)

// Run initializes and runs the game with a bubbletea UI
func Run(appArgs []string, exit chan os.Signal) {
	isDebug := strings.EqualFold(os.Getenv("DEBUG"), "true")

	// Initialize the bubbletea program
	p := tea.NewProgram(
		ui.NewModel(nil, nil, exit), // We'll update these after initialization
		tea.WithAltScreen(),         // Use the full terminal screen
	)

	// Load world and player data
	var w *world.WorldDefinition
	var player actors.Player
	async.InParallel(func() {
		wl := config.Read[world.WorldDefinition]("./data/worlds/dnd_world.yaml", "./data/monsters_dnd.yaml", "./data/weapons.yml")
		w = &wl
		pterm.Success.Println("World loaded.")
	}, func() {
		ps := config.Read[actors.PlayerParams]("./data/test_player.yaml", "./data/weapons.yml")
		player = actors.NewPlayer(ps)
		pterm.Success.Println("Player loaded.")
	})

	// Create the prompt library
	promptLib := ui.NewBubbleTeaPromptLib(p)

	// Initialize game state
	s := state.New(promptLib, player, w.StartRoom, w)
	pterm.Success.Println("Game state initialized.")

	// Create commands
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

	// Create the application
	app := command.NewApp(
		"goblins-and-gold",
		"A daring game of text-based dungeon crawling",
		cmds...,
	)

	pterm.Success.Println("Application configured.")
	pterm.Info.Println("Starting game...")

	// Set up pterm loggers to redirect to bubbletea
	ui.SetupPtermLoggers(p)

	// Start the program in a goroutine
	go func() {
		if _, err := p.Run(); err != nil {
			pterm.Error.Println("Error running program:", err)
			exit <- os.Interrupt
		}
	}()

	// Update the model with the initialized state and app
	p.Send(ui.InitMsg{
		State: s,
		App:   app,
	})
}
