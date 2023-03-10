package path

type PathDefinition struct {
	Room        string `yaml:"room"`
	Description string `yaml:"description"`
}

func (d *PathDefinition) Describe() string {
	return d.Description
}
