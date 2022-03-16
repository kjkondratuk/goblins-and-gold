package equipment

type Weapon struct {
	Name       string
	NoDice     int
	NoSides    int
	DamageType DamageType
	Properties []Property
}

type DamageType int

type Property string

const (
	DT_BLUDGEON = iota
	DT_PIERCING
	DT_SLASHING

	P_LIGHT          = "Light"
	P_HEAVY          = "Heavy"
	P_REACH          = "Reach"
	P_RANGE_20_60    = "Range (20/60)"
	P_RANGE_30_120   = "Range (30/120)"
	P_RANGE_80_320   = "Range (80/320)"
	P_FINESSE        = "Finesse"
	P_VERSATILE_1_8  = "Versatile (1d8)"
	P_VERSATILE_1_10 = "Versatile (1d10)"
	P_TWO_HANDED     = "Two-handed"
	P_AMMUNITION     = "Ammunition"
	P_LOADING        = "Loading"
)

func Shortsword() Weapon {
	return Weapon{
		Name: "Shortword",
		//NoDice:
	}
}
