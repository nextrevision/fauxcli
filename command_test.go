package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCommand(t *testing.T) {
	assert.NotNil(t, validateCommand(Command{}))
	assert.NotNil(t, validateCommand(Command{Name: "test"}))
	assert.NotNil(t, validateCommand(Command{Help: "test"}))
	assert.Nil(t, validateCommand(Command{Name: "test", Help: "test"}))
}

func TestSetCommand(t *testing.T) {
	c := setCommand(Command{
		Name:    "app",
		Help:    "help text",
		Aliases: []string{"app1", "app2"},
	})

	assert.Equal(t, c.Use, "app", "Not equal")
	assert.Equal(t, c.Short, "help text", "Not equal")
	assert.Equal(t, c.Long, "help text", "Not equal")
	assert.Equal(t, c.Aliases, []string{"app1", "app2"}, "Not equal")
}
