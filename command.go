package main

import (
	"html/template"
	"log"
	"os"

	"github.com/Pallinder/go-randomdata"
	"github.com/leekchan/gtf"
	"github.com/spf13/cobra"
)

// Command is the basic command structure
type Command struct {
	Name     string    `yaml:"name"`
	Aliases  []string  `yaml:"aliases"`
	Help     string    `yaml:"help"`
	Output   string    `yaml:"output"`
	Commands []Command `yaml:"commands"`
	Flags    []Flag    `yaml:"flags"`
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
