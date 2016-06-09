package main

import "github.com/spf13/pflag"

// Flag is the basic flag structure
type Flag struct {
	Name    string      `yaml:"name"`
	Short   string      `yaml:"short"`
	Help    string      `yaml:"help"`
	Global  bool        `yaml:"global"`
	Default interface{} `yaml:"default"`
	Type    string      `yaml:"type"`
	value   string      `yaml:"value"`
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
