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
	filename := "cli.yml"
	if fileExists("cli.yaml") {
		filename = "cli.yaml"
	} else if os.Getenv("FAUXCLI_FILE") != "" {
		filename = os.Getenv("FAUXCLI_FILE")
	}

	if !fileExists(filename) && os.Getenv("FAUXCLI_INIT") == "1" {
		if err := initCLIYAML(filename); err != nil {
			log.Fatalf("Could not init fauxcli file %s: %s", filename, err)
		}
	} else if !fileExists(filename) {
		log.Fatalf("%s not found", filename)
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

func initCLIYAML(filename string) error {
	root := Command{
		Name:    "mycliapp",
		Help:    "does something cool",
		Aliases: []string{"myapp", "app"},
		Output:  "Hello, World!\n",
		Flags: []Flag{
			Flag{
				Name:    "debug",
				Short:   "d",
				Help:    "enables debugging",
				Default: false,
				Global:  true,
				Type:    "bool",
			},
		},
		Commands: []Command{
			Command{
				Name:   "subcommand1",
				Help:   "a subcommand",
				Output: "{{ if .Flags.upper.Bool -}}\nHELLO FROM SC1!\n{{ else -}}\nHello from SC1!\n{{ end -}}",
				Flags: []Flag{
					Flag{
						Name:  "upper",
						Short: "u",
						Help:  "converts output to uppercase",
						Type:  "bool",
					},
				},
			},
		},
	}

	data, err := yaml.Marshal(&root)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)
	return err
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
