package mock

import "github.com/stretchr/testify/mock"

type PromptMock struct {
	mock.Mock
}

func (p *PromptMock) Prompt(label string, defaultValue string) (string, error) {
	args := p.Called(label, defaultValue)
	return args.String(0), args.Error(1)
}

func (p *PromptMock) Select(label string, items []string) (int, string, error) {
	args := p.Called(label, items)
	return args.Int(0), args.String(1), args.Error(2)
}
