package interactable

type Action string

type ActionFunc func()

var (
	actions = map[string]ActionFunc{}
)

func DefaultActionSet(a map[string]ActionFunc) {
	actions = a
}

type Interactable struct {
	Description string   `yaml:"description"`
	Actions     []Action `yaml:"actions"`
}

func (i *Interactable) Do(action string) {

}
