package actors

import "github.com/kjkondratuk/goblins-and-gold/stats"

type Combatant interface {
	Dmg(hp int) bool
	BaseStats() stats.BaseStats
	Roll(rollExp string) int
}
