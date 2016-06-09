package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/Pallinder/go-randomdata"
	"github.com/leekchan/gtf"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

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

func setStringFlag(f Flag, fs *pflag.FlagSet) {
	flags[f.Name] = new(string)

	d := ""
	if f.Default != nil {
		d = f.Default.(string)
	}
	fs.StringVarP(
		flags[f.Name].(*string),
		f.Name,
		f.Short,
		d,
		f.Help,
	)
}

func setBoolFlag(f Flag, fs *pflag.FlagSet) {
	flags[f.Name] = new(bool)

	d := false
	if f.Default != nil {
		d = f.Default.(bool)
	}
	fs.BoolVarP(
		flags[f.Name].(*bool),
		f.Name,
		f.Short,
		d,
		f.Help,
	)
}

func setCommand(command Command) *cobra.Command {
	c := &cobra.Command{Use: command.Name}
	if command.Help != "" {
		c.Short = command.Help
		c.Long = command.Help
	}

	if len(command.Aliases) > 0 {
		c.Aliases = command.Aliases
	}

	if command.Output != "" {
		c.Run = func(c *cobra.Command, s []string) {
			funcmap := template.FuncMap{
				"name":      randomdata.SillyName,
				"fullname":  randomFullName,
				"email":     randomdata.Email,
				"city":      randomdata.City,
				"street":    randomdata.Street,
				"address":   randomdata.Address,
				"number":    randomdata.Number,
				"paragraph": randomdata.Paragraph,
				"ipaddress": randomdata.IpV4Address,
				"day":       randomdata.Day,
				"month":     randomdata.Month,
				"date":      randomdata.FullDate,
				"string":    toString,
				"bool":      toBool,
			}
			gtf.Inject(funcmap)

			t, err := template.New("output").Funcs(funcmap).Parse(command.Output)
			if err != nil {
				log.Fatalf("Error parsing template: %s", err.Error())
			}

			err = template.Must(t, err).Execute(os.Stdout, flags)
			if err != nil {
				log.Fatalf("Error parsing template: %s", err.Error())
			}
		}
	}
	return c
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
