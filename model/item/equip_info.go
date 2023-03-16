package item

type EquipInfo struct {
	Slot     string `yaml:"slot"`
	ACBonus  *int   `yaml:"ac_bonus,omitempty"`
	AtkBonus *int   `yaml:"atk_bonus,omitempty"`
	HitBonus *int   `yaml:"hit_bonus,omitempty"`
}
