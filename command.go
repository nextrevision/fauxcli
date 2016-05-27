package main

import (
	"fmt"
	"strings"
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

func processCommand(command Command, flags []Flag, args []string) (Command, []Flag, error) {
	for i, a := range args {
		if strings.HasPrefix(a, "-") {
			flag, err := lookupFlag(a, command.Flags)
			if err != nil {
				return command, flags, err
			}

			flags = append(flags, flag)
		}

		c, err := lookupCommand(a, command.Commands)
		if err != nil {
			return command, flags, err
		}

		return processCommand(c, flags, args[i+1:])
	}

	return command, flags, nil
}

func lookupCommand(arg string, commands []Command) (Command, error) {
	for _, c := range commands {
		if arg == c.Name || contains(c.Aliases, arg) {
			return c, nil
		}
	}

	return Command{}, fmt.Errorf("No such command: %s", arg)
}
