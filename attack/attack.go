package attack

type AttackSet map[string]Attack

type Attack struct {
	Bonus  int `yaml:"bonus"`
	Range  int
	Damage []AttackDamage
}

type AttackDamage struct {
	Roll       string
	Bonus      int
	DamageType Type
}

type MultiAttackSet []MultiAttack

type MultiAttack struct {
	Number  int
	Attacks []Attack
}
