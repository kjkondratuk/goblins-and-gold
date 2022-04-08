package room

type Room struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Pathways    []string `yaml:"pathways"`
}
