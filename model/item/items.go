package item

type Type string

type Item struct {
	Type        Type   `yaml:"type"`
	Description string `yaml:"description"`
	Quantity    int    `yaml:"quantity"`
	Unit        string `yaml:"unit"`
}

func (i Item) Describe() string {
	return i.Description
}
