package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFlag(t *testing.T) {
	flags := map[string]Flag{}
	flags["flag1"] = Flag{
		Name:  "flag1",
		Help:  "help text",
		Short: "f",
	}

	assert.NotNil(t, validateFlag(Flag{}, flags))
	assert.NotNil(t, validateFlag(Flag{Name: "test"}, flags))
	assert.NotNil(t, validateFlag(Flag{Help: "test"}, flags))
	assert.NotNil(t, validateFlag(Flag{Name: "flag1", Help: "h"}, flags))
	assert.NotNil(t, validateFlag(Flag{Name: "flag2", Help: "h", Short: "f"}, flags))
	assert.Nil(t, validateFlag(Flag{Name: "test", Help: "test"}, flags))
}

func TestSetStringFlag(t *testing.T) {
	sVal := "val"
	bVal := true
	s := Flag{Name: "flag", Help: "h", Short: "f", value: &sVal, Type: "string"}
	b := Flag{Name: "flag", Help: "h", Short: "f", value: &bVal, Type: "bool"}

	assert.Equal(t, s.String(), "val", "Should be equal")
	assert.Equal(t, s.Value().(*string), &sVal, "Should be equal")
	assert.Equal(t, b.Bool(), true, "Should be equal")
	assert.Equal(t, b.Value().(*bool), &bVal, "Should be equal")
}
