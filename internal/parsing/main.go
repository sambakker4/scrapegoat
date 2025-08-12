package parsing

import (
	"strings"
)

type Ordering int

const (
	ASC Ordering = iota
	DESC
	NONE
)

var keywords = []string{"select", "from", "limit", "where", "order", "by", "distinct"}

type CliConfig struct {
	Select   []string
	From     string
	Limit    int
	Where    string
	OrderBy  Ordering
	Distinct bool
}


func ParseCLI(text string) (*CliConfig, error) {
	query := CleanInput(text)
	config := &CliConfig{}

	query, err := parseSelect(query, config)
	if err != nil {
		return &CliConfig{}, err
	}	

	return config, nil
}

func CleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	return strings.Fields(lowercaseText)
}
