package attack

const (
	Bludgeoning = "Bludgeoning"
	Piercing    = "Piercing"
	Slashing    = "Slashing"
	Thunder     = "Thunder"
	Force       = "Force"
	Psychic     = "Psychic"
	Fire        = "Fire"
	Lightning   = "Lightning"
)

type Type string

type Damage struct {
	Amount     int  `yaml:"amount"`
	DamageType Type `yaml:"damage_type"`
}
