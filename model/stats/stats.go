package stats

const (
	StatNameLevel        = "Lvl"
	StatNameStrength     = "Str"
	StatNameDexterity    = "Dex"
	StatNameConstitution = "Con"
	StatNameIntelligence = "Int"
	StatNameWisdom       = "Wis"
	StatNameCharisma     = "Cha"
)

type BaseStats struct {
	Lvl int `yaml:"lvl"`
	Str int `yaml:"str"`
	Dex int `yaml:"dex"`
	Con int `yaml:"con"`
	Int int `yaml:"int"`
	Wis int `yaml:"wis"`
	Cha int `yaml:"cha"`
}

func (s BaseStats) GetByName(name string) (int, bool) {
	switch name {
	case StatNameStrength:
		return s.Str, true
	case StatNameDexterity:
		return s.Dex, true
	case StatNameConstitution:
		return s.Con, true
	case StatNameIntelligence:
		return s.Int, true
	case StatNameWisdom:
		return s.Wis, true
	case StatNameCharisma:
		return s.Cha, true
	case StatNameLevel:
		return s.Lvl, true
	default:
		return 0, false
	}
}

func (s BaseStats) ModifierByName(name string) (int, bool) {
	if mod, ok := s.GetByName(name); ok {
		return calcModifier(mod), true
	}
	return 0, false
}

func (s BaseStats) DexMod() int {
	return calcModifier(s.Dex)
}

func calcModifier(value int) int {
	v := (value - 10) / 2
	return v
}
