package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

var flags = map[string]Flag{}

func main() {
	filename := "cli.yaml"
	if os.Getenv("CLIMOCK_FILE") != "" {
		filename = os.Getenv("CLIMOCK_FILE")
	}

	cli, err := loadCLIYAML(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = processCommands(cli).Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func processCommands(command Command) *cobra.Command {
	cobraCommand := setCommand(command)
	if len(command.Commands) > 0 {
		for _, c := range command.Commands {
			err := validateCommand(c)
			if err != nil {
				log.Fatalln(err.Error())
			}

			cobraCommand.AddCommand(processCommands(c))
		}
	}

	if len(command.Flags) > 0 {
		for _, f := range command.Flags {
			err := validateFlag(f, flags)
			if err != nil {
				log.Fatalln(err.Error())
			}

			flagSet := cobraCommand.Flags()

			if f.Global {
				flagSet = cobraCommand.PersistentFlags()
			}

			switch f.Type {
			case "string":
				setStringFlag(f, flagSet)
			case "bool":
				setBoolFlag(f, flagSet)
			case "int":
				setIntFlag(f, flagSet)
			case "float":
				setFloatFlag(f, flagSet)
			default:
				setStringFlag(f, flagSet)
			}
		}
	}

	return cobraCommand
}

func loadCLIYAML(filename string) (Command, error) {
	command := Command{}
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return command, fmt.Errorf("cli.yml not found in current directory")
	}

	err = yaml.Unmarshal(contents, &command)
	if err != nil {
		return command, err
	}

	return command, nil
}
