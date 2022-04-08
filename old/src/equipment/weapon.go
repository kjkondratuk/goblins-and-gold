// Package equipment provides necessary weapon and non-combat item detail implementations to parse and handle interactions
// with them.
package equipment

type Weapon struct {
	Name       string
	NoDice     int
	NoSides    int
	DamageType DamageType
	Properties []Property
}

type Property string

const (
	// P_LIGHT : a descriptor for a weapon that can be dual-wielded
	P_LIGHT = "Light"

	// P_HEAVY : a descriptor for a weapon that is difficult for a small creature to wield
	P_HEAVY = "Heavy"

	// P_REACH : a descriptor for a weapon that has melee range beyond the typical 5'
	P_REACH = "Reach"

	// P_RANGE_20_60 : a descriptor for a ranged weapon with 20-60' of range
	P_RANGE_20_60 = "Range (20/60)"

	// P_RANGE_30_120 : a descriptor for a ranged weapon with 30-120' of range
	P_RANGE_30_120 = "Range (30/120)"

	// P_RANGE_80_320 : a descriptor for a ranged weapon with 80-320' of range
	P_RANGE_80_320 = "Range (80/320)"

	// P_FINESSE : a descriptor for a weapon which can use either strength or dexterity modifier
	P_FINESSE = "Finesse"

	// P_VERSATILE_1_8 : a descriptor for a weapon that can be wielded with two hands and deal 1d8 damage
	P_VERSATILE_1_8 = "Versatile (1d8)"

	// P_VERSATILE_1_10 : a descriptor for a weapon that can be wielded with two hands and deal 1d10 damage
	P_VERSATILE_1_10 = "Versatile (1d10)"

	// P_TWO_HANDED : a descriptor for a weapon that requires two hands to perform attacks
	P_TWO_HANDED = "Two-handed"

	// P_AMMUNITION : a descriptor indicating that a weapon requires a player to possess ammunition to fire
	P_AMMUNITION = "Ammunition"

	// P_LOADING : a descriptor for a weapon that requires the use of a bonus action to reload
	P_LOADING = "Loading"
)

// TODO : probably want to parse weapons from a file or something using a standard format instead of hard-coding accessors.

func Shortsword() Weapon {
	return Weapon{
		Name: "Shortword",
		//NoDice:
	}
}
