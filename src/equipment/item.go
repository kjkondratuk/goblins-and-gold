package equipment

// DamageType : The type of damage a given weapon/item applies
type DamageType int

const (
	// DT_BLUDGEON : Damage type indicating bludgeoning type damage
	DT_BLUDGEON = iota

	// DT_PIERCING : Damage type indicating piercing type damage
	DT_PIERCING

	// DT_SLASHING : Damage type indicating slashing type damage
	DT_SLASHING
)
