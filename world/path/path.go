package path

type Definition struct {
	Room        string `yaml:"room"`
	Description string `yaml:"description"`
}

func (d *Definition) Describe() string {
	return d.Description
}
