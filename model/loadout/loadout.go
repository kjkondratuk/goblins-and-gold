package loadout

import "github.com/kjkondratuk/goblins-and-gold/model/item"

const (
	SlotHelm   = "Helm"
	SlotNeck   = "Neck"
	SlotTorso  = "Torso"
	SlotHands  = "Hands"
	SlotWrists = "Wrists"
	SlotRings  = "Rings"
	SlotFeet   = "Feet"
)

type Loadout struct {
	Helm   *item.Item    `yaml:"helm,omitempty"`
	Neck   *item.Item    `yaml:"neck,omitempty"`
	Torso  *item.Item    `yaml:"torso,omitempty"`
	Hands  *item.Item    `yaml:"hands,omitempty"`
	Wrists [2]*item.Item `yaml:"wrists,omitempty"`
	Rings  [2]*item.Item `yaml:"rings,omitempty"`
	Feet   *item.Item    `yaml:"feet,omitempty"`
}

func (l *Loadout) Equip(i *item.Item, idx int) bool {
	if i != nil && i.EquipInfo != nil {
		switch i.EquipInfo.Slot {
		case SlotHelm:
			l.Helm = i
		case SlotNeck:
			l.Neck = i
		case SlotTorso:
			l.Torso = i
		case SlotHands:
			l.Hands = i
		case SlotWrists:
			if idx == 1 || idx == 0 {
				l.Wrists[idx] = i
			} else {
				return false
			}
		case SlotRings:
			if idx == 1 || idx == 0 {
				l.Rings[idx] = i
			} else {
				return false
			}
		case SlotFeet:
			l.Feet = i
		default:
			return false
		}
	} else {
		return false
	}
	return true
}

func (l *Loadout) Unequip(slot string, idx int) bool {
	switch slot {
	case "Helm":
		l.Helm = nil
	case "Neck":
		l.Neck = nil
	case "Torso":
		l.Torso = nil
	case "Hands":
		l.Hands = nil
	case "Wrists":
		if idx == 1 || idx == 0 {
			l.Wrists[idx] = nil
		} else {
			return false
		}
	case "Rings":
		if idx == 1 || idx == 0 {
			l.Rings[idx] = nil
		} else {
			return false
		}
	case "Feet":
		l.Feet = nil
	default:
		return false
	}
	return true
}
