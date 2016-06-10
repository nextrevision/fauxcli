package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Flag is the basic flag structure
type Flag struct {
	Name    string      `yaml:"name"`
	Short   string      `yaml:"short"`
	Help    string      `yaml:"help"`
	Global  bool        `yaml:"global"`
	Default interface{} `yaml:"default"`
	Type    string      `yaml:"type"`
	value   interface{}
}

func validateFlag(flag Flag, flags map[string]Flag) error {
	if flag.Name == "" {
		return fmt.Errorf("Missing name for flag: %+v", flag)
	}

	if flag.Help == "" {
		return fmt.Errorf("Missing help for flag: %s", flag.Name)
	}

	for _, f := range flags {
		if f.Name == flag.Name {
			return fmt.Errorf("Duplicate flag found for %s", flag.Name)
		}

		if flag.Short != "" && f.Short == flag.Short {
			return fmt.Errorf("Duplicate flag short name found for %s", flag.Name)
		}
	}

	return nil
}

func setStringFlag(f Flag, fs *pflag.FlagSet) {
	f.value = new(string)
	flags[f.Name] = f

	d := ""
	if f.Default != nil {
		d = f.Default.(string)
	}

	fs.StringVarP(
		flags[f.Name].value.(*string),
		f.Name,
		f.Short,
		d,
		f.Help,
	)
}

func setBoolFlag(f Flag, fs *pflag.FlagSet) {
	f.value = new(bool)
	flags[f.Name] = f

	d := false
	if f.Default != nil {
		d = f.Default.(bool)
	}

	fs.BoolVarP(
		flags[f.Name].value.(*bool),
		f.Name,
		f.Short,
		d,
		f.Help,
	)
}

func setIntFlag(f Flag, fs *pflag.FlagSet) {
	f.value = new(int)
	flags[f.Name] = f

	d := 0
	if f.Default != nil {
		d = f.Default.(int)
	}

	fs.IntVarP(
		flags[f.Name].value.(*int),
		f.Name,
		f.Short,
		d,
		f.Help,
	)
}

func setFloatFlag(f Flag, fs *pflag.FlagSet) {
	f.value = new(float64)
	flags[f.Name] = f

	d := 0.0
	if f.Default != nil {
		d = f.Default.(float64)
	}

	fs.Float64VarP(
		flags[f.Name].value.(*float64),
		f.Name,
		f.Short,
		d,
		f.Help,
	)
}

// Value returns the generic interface value of a flag
func (f Flag) Value() interface{} {
	return f.value
}

// String returns the string value of a flag
func (f Flag) String() string {
	if f.Type == "string" || f.Type == "" {
		v := f.value.(*string)
		return *v
	}

	return ""
}

// Bool returns the string value of a flag
func (f Flag) Bool() bool {
	if f.Type == "bool" {
		v := f.value.(*bool)
		return *v
	}

	return false
}

// Int returns the int value of a flag
func (f Flag) Int() int {
	if f.Type == "int" {
		v := f.value.(*int)
		return *v
	}

	return 0
}

// Float returns the int value of a flag
func (f Flag) Float() float64 {
	if f.Type == "float" {
		v := f.value.(*float64)
		return *v
	}

	return 0
}
