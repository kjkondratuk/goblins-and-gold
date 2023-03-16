package command

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

type DefaultCommandCompleter struct{}

func (c DefaultCommandCompleter) Completer(command Command) prompt.Completer {
	return func(d prompt.Document) []prompt.Suggest {
		parts := strings.Split(d.TextBeforeCursor(), " ")
		var foundCommand Command
		foundCommand = command
		var parentCommand Command
		for i, cmd := range parts {
			if i == len(parts)-1 {
				parentCommand = foundCommand
			}
			foundCommand = getChildByName(foundCommand, cmd)
		}

		var suggestions []prompt.Suggest
		if parentCommand != nil {
			for _, pc := range parentCommand.Subcommands() {
				if strings.HasPrefix(pc.Name(), d.GetWordBeforeCursor()) {
					suggestions = append(suggestions, prompt.Suggest{
						Text:        pc.Name(),
						Description: pc.Description(),
					})
				}
			}
		}

		return suggestions
	}
}

func getChildByName(cmd Command, name string) Command {
	if len(name) <= 0 {
		return cmd
	}
	if cmd != nil {
		for _, c := range cmd.Subcommands() {
			hasAlias := false
			for _, a := range c.Aliases() {
				if a == name {
					hasAlias = true
				}
			}
			if c.Name() == name || hasAlias {
				return c
			}
		}
	}
	return nil
}
