package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApp(t *testing.T) {
	t.Run("should create app with help", func(t *testing.T) {
		a := NewApp("app name", "an application")

		assert.NotNil(t, a)
		assert.NotEmpty(t, a.Subcommands())
		assert.Len(t, a.Subcommands(), 1)
		assert.Equal(t, "help", a.Subcommands()[0].Name())
	})

	t.Run("should create an app with commads", func(t *testing.T) {
		com1 := &baseCommand{
			name:        "",
			description: "",
			usage:       "",
			aliases:     nil,
			subcommands: nil,
		}
		com2 := &baseCommand{
			name:        "",
			description: "",
			usage:       "",
			aliases:     nil,
			subcommands: nil,
		}

		a := NewApp("app name", "an application", com1, com2)

		assert.NotNil(t, a)
		assert.NotEmpty(t, a.Subcommands())
		assert.Len(t, a.Subcommands(), 3)
		assert.Equal(t, com1, a.Subcommands()[0])
		assert.Equal(t, com2, a.Subcommands()[1])
		assert.Equal(t, "help", a.Subcommands()[2].Name())
	})
}

func Test_app_Run(t *testing.T) {
	type fields struct {
		baseCommand baseCommand
	}
	type args struct {
		s    state.State
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"should not error when run without commands or usage",
			fields{baseCommand: baseCommand{
				name:        "",
				description: "",
				usage:       "",
				aliases:     nil,
				subcommands: nil,
			}},
			args{
				s:    nil,
				args: nil,
			},
			false,
		}, {
			"should not error when run without commands but with usage",
			fields{baseCommand: baseCommand{
				name:        "",
				description: "",
				usage:       "example of usage",
				aliases:     nil,
				subcommands: nil,
			}},
			args{
				s:    nil,
				args: nil,
			},
			false,
		}, {
			"should not error when run with subcommand",
			fields{baseCommand: baseCommand{
				name:        "",
				description: "",
				usage:       "example of usage",
				aliases:     nil,
				subcommands: []Command{
					NewHelpCommand(NewApp("", "")),
				},
			}},
			args{
				s:    nil,
				args: nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &app{
				baseCommand: tt.fields.baseCommand,
			}
			if err := a.Run(tt.args.s, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
