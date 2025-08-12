package main

import (
	"testing"
	"github.com/sambakker4/scrapegoat/internal/parsing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCleanInput(t *testing.T) {
	// Standard Input
	input := "asdf ASDF asdf"
	expected := []string{"asdf", "asdf", "asdf"}
	cleanedInput := parsing.CleanInput(input)
	assert.Equal(t, cleanedInput, expected)

	// Extra Whitespace
	input = "   hey         hi              --help"
	expected = []string{"hey", "hi", "--help"}
	cleanedInput = parsing.CleanInput(input)
	assert.Equal(t, cleanedInput, expected)

	// All Caps
	input = "HEY    THERE   HI"
	expected = []string{"hey", "there", "hi"}
	cleanedInput = parsing.CleanInput(input)
	assert.Equal(t, cleanedInput, expected)
}

func TestIsValidURL(t *testing.T) {
	// Standard Input
	input := "http://google.com"
	output := isValidURL(input)
	require.True(t, output)

	// Invalid URL
	input = "asdfadsfasdf"
	output = isValidURL(input)
	require.False(t, output)
}
