package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	cli, err := loadCLIYAML()
	if err != nil {
		log.Fatal(err)
	}

	command, flags, err := processCommand(cli, []Flag{}, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if len(command.Commands) > 0 {
		printHelp(cli, command)
	} else {
		printOutput(command.Output, flags)
	}
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printHelp(cli Command, command Command) {
	usage := fmt.Sprintf("usage: %s", cli.Name)
	if len(cli.Flags) > 0 {
		usage = fmt.Sprintf("%s [flags]", usage)
	}
	if cli.Name != command.Name {
		usage = fmt.Sprintf("%s %s", usage, command.Name)
	}
	if len(cli.Commands) > 0 {
		usage = fmt.Sprintf("%s command", usage)
	}
	fmt.Printf("%s\n", usage)

	if len(command.Commands) > 0 {
		fmt.Printf("\nCommands:\n")
		for _, c := range command.Commands {
			fmt.Printf("  %-20s %s\n", c.Name, c.Help)
		}
	}

	if len(command.Flags) > 0 {
		fmt.Printf("\nFlags:\n")
		for _, f := range command.Flags {
			flagString := fmt.Sprintf("--%s", f.Name)
			if f.Short != "" {
				flagString = fmt.Sprintf("-%s|%s", f.Short, flagString)
			}
			fmt.Printf("  %-20s %s\n", flagString, f.Help)
		}
	}

	if len(cli.Flags) > 0 {
		fmt.Printf("\nGlobal Flags:\n")
		for _, f := range cli.Flags {
			flagString := fmt.Sprintf("--%s", f.Name)
			if f.Short != "" {
				flagString = fmt.Sprintf("-%s|%s", f.Short, flagString)
			}
			fmt.Printf("  %-20s %s\n", flagString, f.Help)
		}
	}
}

func printOutput(output string, flags []Flag) error {
	funcmap := template.FuncMap{
		"namevar": namevar,
	}
	data := struct {
		Flags []Flag
	}{
		Flags: flags,
	}
	t := template.Must(template.New("output").Funcs(funcmap).Parse(output))
	err := t.Execute(os.Stdout, data)
	return err
}
