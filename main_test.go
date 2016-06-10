package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessCommands(t *testing.T) {
	c := processCommands(Command{
		Name:    "mycliapp",
		Help:    "does something cool",
		Aliases: []string{"myapp", "app"},
		Output:  "Hello World!\n",
		Flags: []Flag{
			Flag{
				Name:    "debug",
				Help:    "enables debugging",
				Short:   "d",
				Default: false,
				Global:  true,
				Type:    "bool",
			},
		},
		Commands: []Command{
			Command{
				Name: "subcommand1",
				Help: "a subcommand",
				Flags: []Flag{
					Flag{
						Name:  "upper",
						Help:  "converts output to uppercase",
						Short: "u",
						Type:  "bool",
					},
				},
				Output: "{{ if .Flags.upper.Bool -}}\nHELLO FROM SC1!\n{{ else -}}\nHello from SC1!\n{{ end -}}\n",
			},
		},
	})

	assert.Equal(t, c.Use, "mycliapp", "Use should be set to mycliapp")
	assert.Equal(t, c.Short, "does something cool", "Short is not equal")
	assert.Equal(t, c.Long, "does something cool", "Long is not equal")
	assert.Equal(t, c.Aliases, []string{"myapp", "app"}, "Aliases is not equal")
	assert.Equal(t, len(flags), 2, "Flags length is not equal")

	commands := c.Commands()
	assert.Equal(t, len(commands), 1, "Commands length is not equal")
	assert.Equal(t, commands[0].Use, "subcommand1", "Not equal")
	assert.Equal(t, commands[0].Short, "a subcommand", "Not equal")
	assert.Equal(t, commands[0].Long, "a subcommand", "Not equal")

	assert.Nil(t, c.Execute())
}

func TestLoadCLIYAML(t *testing.T) {
	want := Command{
		Name:    "mycliapp",
		Help:    "does something cool",
		Aliases: []string{"myapp", "app"},
		Output:  "Hello, World!\n",
		Flags: []Flag{
			Flag{
				Name:    "debug",
				Help:    "enables debugging",
				Short:   "d",
				Default: false,
				Global:  true,
				Type:    "bool",
			},
		},
		Commands: []Command{
			Command{
				Name: "subcommand1",
				Help: "a subcommand",
				Flags: []Flag{
					Flag{
						Name:  "upper",
						Help:  "converts output to uppercase",
						Short: "u",
						Type:  "bool",
					},
				},
				Output: "{{ if .Flags.upper.Bool -}}\nHELLO FROM SC1!\n{{ else -}}\nHello from SC1!\n{{ end -}}\n",
			},
			Command{
				Name: "subcommand2",
				Help: "another subcommand with children",
				Commands: []Command{
					Command{
						Name:   "child1",
						Help:   "the first child command",
						Output: "Hello from child1\n",
					},
					Command{
						Name:   "child2",
						Help:   "the second child command",
						Output: "Hello from child2\n",
					},
				},
			},
		},
	}

	c, err := loadCLIYAML("examples/mycliapp.yaml")

	assert.Nil(t, err)

	if !reflect.DeepEqual(c, want) {
		t.Errorf("Loaded data not equal, data returned %+v, want %+v", c, want)
	}
}
