package item

type Type string

type Item struct {
	Type        Type       `yaml:"type"`
	EquipInfo   *EquipInfo `yaml:"equip_info"`
	Description string     `yaml:"description"`
	Quantity    int        `yaml:"quantity"`
	Unit        string     `yaml:"unit"`
}

func (i Item) Describe() string {
	return i.Description
}
