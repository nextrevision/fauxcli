package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

var flags = map[string]interface{}{}

func main() {
	cli, err := loadCLIYAML()
	if err != nil {
		log.Fatal(err)
	}

	processCommands(cli).Execute()
}

func processCommands(command Command) *cobra.Command {
	root := setCommand(command)
	if len(command.Commands) > 0 {
		for _, c := range command.Commands {
			root.AddCommand(processCommands(c))
		}
	}

	if len(command.Flags) > 0 {
		for _, f := range command.Flags {
			flagSet := root.Flags()

			if f.Global {
				flagSet = root.PersistentFlags()
			}

			switch f.Type {
			case "string":
				setStringFlag(f, flagSet)
			case "bool":
				setBoolFlag(f, flagSet)
			default:
				setStringFlag(f, flagSet)
			}
		}
	}

	return root
}

func loadCLIYAML() (Command, error) {
	command := Command{}
	contents, err := ioutil.ReadFile("cli.yml")
	if err != nil {
		return command, fmt.Errorf("cli.yml not found in current directory")
	}

	err = yaml.Unmarshal(contents, &command)
	if err != nil {
		return command, err
	}

	return command, nil
}
