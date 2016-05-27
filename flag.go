package main

import (
	"fmt"
	"strings"
)

// Flag is the basic flag structure
type Flag struct {
	Name     string `yaml:"name"`
	Short    string `yaml:"short"`
	Required bool   `yaml:"required"`
	Help     string `yaml:"help"`
	value    string `yaml:"value"`
}

func lookupFlag(a string, flags []Flag) (Flag, error) {
	arg := strings.TrimLeft(a, "-")
	for _, f := range flags {
		if arg == f.Name || arg == strings.TrimLeft(f.Short, "-") {
			return f, nil
		}
	}

	return Flag{}, fmt.Errorf("No such flag: %s", a)
}
