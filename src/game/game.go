// Package game is responsible for constructing, persisting, and coordinating high level interactions with the game state
package game

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/src/navigator"
	"github.com/kjkondratuk/goblins-and-gold/src/player"
	"github.com/kjkondratuk/goblins-and-gold/src/room"
	"github.com/kjkondratuk/goblins-and-gold/src/world"
	"github.com/olekukonko/ts"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	worldFile  = "./config/test_world.yaml"
	playerFile = "./config/test_player.yaml"
)

func Start() {
	// setup exit listener
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	w := readConfig[world.WorldData](worldFile)
	p := readConfig[player.PlayerData](playerFile)

	startRoom := room.NewRoom(room.WithRoomData(w.StartRoom))
	/*nav :=*/ navigator.NewNavigatorFrom(player.NewPlayer(player.WithPlayerData(p)), startRoom)

	fmt.Printf("World: %+v\n", w)
	fmt.Printf("Player: %+v\n", p)

	fmt.Printf("Game Client Initialized\n")

	go func() {
		// TODO : do I really want to check terminal size in this way?  Maybe we shouldn't be updating the terminal buffer...
		size, err := ts.GetSize()
		if err != nil {
			log.Fatalf("Could not get terminal size: %+v", err)
		}
		fmt.Printf("width: %d - height: %d\n", size.Col(), size.Row())
		select {}
	}()

	sig := <-exit
	fmt.Printf("%s received, exiting...\n", sig)
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
