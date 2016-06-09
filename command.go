package main

// Command is the basic command structure
type Command struct {
	Name     string    `yaml:"name"`
	Aliases  []string  `yaml:"aliases"`
	Help     string    `yaml:"help"`
	Output   string    `yaml:"output"`
	Commands []Command `yaml:"commands"`
	Flags    []Flag    `yaml:"flags"`
}
