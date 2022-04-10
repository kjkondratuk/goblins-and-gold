package player

import (
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
)

type Player struct {
	// TODO : this should probably be moved into BaseStats since it's dependent upon class and Con
	HP        int             `yaml:"hp"`
	BaseStats stats.BaseStats `yaml:"stats"`
	Inventory []item.Item     `yaml:"inventory"`
}

// Dmg : applies damage to a player's hitpoints.
// Parameters:
//   - hp - int - the number of hitpoints to remove
// Returns:
//   - bool - whether or not the damage could be applied
func (p Player) Dmg(hp int) bool {
	if hp > 0 {
		p.HP -= hp
		if p.HP < 0 {
			p.HP = 0
		}
		return true
	}
	return false
}
