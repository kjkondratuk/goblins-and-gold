package attack

type AttackSet map[string]Attack

type Attack struct {
	Bonus  int      `yaml:"bonus"`
	Range  int      `yaml:"range"`
	Damage []Damage `yaml:"damage"`
}

type Damage struct {
	Roll       string `yaml:"roll"`
	Bonus      int    `yaml:"bonus"`
	DamageType Type   `yaml:"damage_type"`
}

type MultiAttack struct {
	Number  int      `yaml:"number"`
	Attacks []Attack `yaml:"attacks"`
}
