package parsing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCli(t *testing.T) {

}

func TestCleanInput(t *testing.T) {
	// Standard Input
	input := "asdf ASDF asdf"
	expected := []string{"asdf", "asdf", "asdf"}
	cleanedInput := CleanInput(input)
	assert.Equal(t, cleanedInput, expected)

	// Extra Whitespace
	input = "   hey         hi              --help"
	expected = []string{"hey", "hi", "--help"}
	cleanedInput = CleanInput(input)
	assert.Equal(t, cleanedInput, expected)

	// All Caps
	input = "HEY    THERE   HI"
	expected = []string{"hey", "there", "hi"}
	cleanedInput = CleanInput(input)
	assert.Equal(t, cleanedInput, expected)
}

func TestParseSelect(t *testing.T) {
	// Standard Input
	config := &CliConfig{}
	expectedConfig := &CliConfig{
		Select: []string{"asdfasd", "asdfasdfads", "asdfsd"},
	}
	input := []string{"select", "asdfasd,", "asdfasdfads,", "asdfsd"}
	expected := []string(nil)
	output, err := parseSelect(input, config)
	require.NoError(t, err)
	assert.Equal(t, expected, output)
	assert.Equal(t, expectedConfig, config)

	// Other statements 
	config = &CliConfig{}
	expectedConfig = &CliConfig{
		Select: []string{"asdfasd", "asdfasdfads", "asdfsd"},
	}
	input = []string{"select", "asdfasd,", "asdfasdfads,", "asdfsd", "limit", "5"}
	expected = []string{"limit", "5"}
	output, err = parseSelect(input, config)
	require.NoError(t, err)
	assert.Equal(t, expected, output)
	assert.Equal(t, expectedConfig, config)

	// Lots of other statements
	config = &CliConfig{}
	expectedConfig = &CliConfig{
		Select: []string{"asdfasd", "asdfasdfads", "asdfsd"},
	}
	input = []string{"limit", "6", "select", "asdfasd,", "asdfasdfads,", "asdfsd", "where", "id", "=", "7", "distinct"}
	expected = []string{"limit", "6", "where", "id", "=", "7", "distinct"}
	output, err = parseSelect(input, config)
	require.NoError(t, err)
	assert.Equal(t, expected, output)
	assert.Equal(t, expectedConfig, config)
}
