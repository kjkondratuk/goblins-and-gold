package config

import (
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/player"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"io/ioutil"
	"log"
)

type configDataType interface {
	world.World | player.PlayerData
}

func Read[T configDataType](f string) T {
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
