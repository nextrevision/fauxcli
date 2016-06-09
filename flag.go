package main

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
