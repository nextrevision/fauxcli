package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"gopkg.in/yaml.v2"
)

var flags = map[string]Flag{}

func main() {
	var filename string
	if fileExists("cli.yaml") {
		filename = "cli.yaml"
	} else if fileExists("cli.yml") {
		filename = "cli.yml"
	} else if fileExists(os.Getenv("CLIMOCK_FILE")) {
		filename = os.Getenv("CLIMOCK_FILE")
	} else {
		log.Fatalf("cli.yml not found")
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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func loadCLIYAML(filename string) (Command, error) {
	command := Command{}
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return command, err
	}

	err = yaml.Unmarshal(contents, &command)
	if err != nil {
		return command, err
	}

	return command, nil
}
