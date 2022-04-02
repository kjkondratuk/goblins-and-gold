// Package game is responsible for constructing, persisting, and coordinating high level interactions with the game state
package game

import (
	"bytes"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/src/navigator"
	"github.com/kjkondratuk/goblins-and-gold/src/player"
	"github.com/kjkondratuk/goblins-and-gold/src/room"
	"github.com/kjkondratuk/goblins-and-gold/src/world"
	"github.com/pterm/pterm"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	worldFile  = "./config/test_world.yaml"
	playerFile = "./config/test_player.yaml"
)

var (
	completer = readline.NewPrefixCompleter(
		readline.PcItem("help - print this message"),
		readline.PcItem("quit - exit the game"),
		readline.PcItem("look - examine your surroundings", readline.PcItemDynamic(func(item string) []string {
			return []string{
				"<some item> - examine a specific item in the room",
			}
		})),
	)
)

func Start() {
	// setup exit listener
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	w := readConfig[world.WorldData](worldFile)
	p := readConfig[player.PlayerData](playerFile)

	startRoom := room.NewRoom(room.WithRoomData(w.StartRoom))
	nav := navigator.NewNavigatorFrom(player.NewPlayer(player.WithPlayerData(p)), startRoom)

	fmt.Printf("World: %+v\n", w)
	fmt.Printf("Player: %+v\n", p)

	fmt.Printf("Game Client Initialized\n")

	//go func() {
	reader, err := readline.NewEx(&readline.Config{
		Prompt:          pterm.LightYellow(" > "),
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
		FuncFilterInputRune: func(r rune) (rune, bool) {
			switch r {
			// block CtrlZ feature
			case readline.CharCtrlZ:
				return r, false
			}
			return r, true
		},
	})
	if err != nil {
		log.Fatalf("Could not initiate REPL.  Exiting.")
	}
	defer reader.Close()

	pterm.DefaultHeader.WithFullWidth().Println("goblins-and-gold")

mainLoop:
	for {
		line, err := reader.Readline()
		// handle error/exit conditions
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		// clean up leading/trailing spaces
		line = strings.TrimSpace(line)

		// process commands
		switch {
		case strings.HasPrefix(line, "look"):
			tok := strings.Split(line, " ")
			if len(tok) == 1 {
				fmt.Println(nav.Look())
			} else if len(tok) == 2 {
				switch tok[1] {
				case "help":
					fmt.Println(completer.Tree("look"))
				}
			}
		case strings.HasPrefix(line, "help"):
			usage(0, reader.Stdout())
		case strings.HasPrefix(line, "quit"):
			exit <- syscall.SIGINT
			break mainLoop
		default:
			fmt.Printf(pterm.Red(fmt.Sprintf("\"%s\" is not a valid command.\n", line)))
			usage(0, reader.Stdout())
		}
	}

	sig := <-exit
	fmt.Printf("%s received, exiting...\n", sig)
}

func usage(level int, w io.Writer) {
	_, _ = io.WriteString(w, "Usage:\n")
	buf := bytes.NewBufferString("")
	completer.Print("    ", level, buf)
	_, _ = io.WriteString(w, buf.String())
}

type configDataType interface {
	world.WorldData | player.PlayerData
}

func readConfig[T configDataType](f string) T {
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("error reading file: %v ", err)
	}

	// If we wind up needing more documents, we might loop like:
	// https://stackoverflow.com/questions/70920334/parse-yaml-files-with-in-it
	var c T
	if err = yaml.UnmarshalWithOptions(yamlFile, &c, yaml.ReferenceFiles(f)); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
