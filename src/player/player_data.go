package player

import "github.com/kjkondratuk/goblins-and-gold/src/stats"

type PlayerData struct {
	// TODO : this should probably be moved into BaseStats since it's dependent upon class and Con
	HP        int                `yaml:"hp"`
	BaseStats stats.BaseStatData `yaml:"stats"`
}
