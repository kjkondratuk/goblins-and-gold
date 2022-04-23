package actors

import "github.com/kjkondratuk/goblins-and-gold/stats"

type Combatant interface {
	Name() string
	Health() int
	Dmg(hp int) bool
	BaseStats() stats.BaseStats
	Roll(rollExp string) int
	Attack(c Combatant)
}
