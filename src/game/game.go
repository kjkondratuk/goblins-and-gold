// Package game is responsible for constructing, persisting, and coordinating high level interactions with the game state
package game

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/src/world"
	"io/ioutil"
	"log"
	"os"
)

func Start() {
	yamlFile, err := ioutil.ReadFile("config/test_world.yaml")
	if err != nil {
		log.Fatalf("error reading world file: %v ", err)
	}

	var w world.WorldData
	err = yaml.UnmarshalWithOptions(yamlFile, &w, yaml.ReferenceFiles("config/test_world.yaml"))
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Printf("%+v", w)
	os.Exit(0)
}
