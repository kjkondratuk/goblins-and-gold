package ux

import (
	"github.com/manifoldco/promptui"
)

type promptLib struct{}

//go:generate mockery --name PromptLib
type PromptLib interface {
	Prompt(label string, defaultValue string) (string, error)
	Select(label string, items []string) (int, string, error)
}

func NewPromptUiLib() PromptLib {
	return &promptLib{}
}

// Prompt is a wrapper function for promptui.Prompt
func (pl *promptLib) Prompt(label string, defaultValue string) (string, error) {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

// Select is a wrapper function for promptui.Select
func (pl *promptLib) Select(label string, items []string) (int, string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	idx, result, err := prompt.Run()

	if err != nil {
		return -1, "", err
	}

	return idx, result, nil
}
