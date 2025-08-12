package parsing

import (
	"errors"
	"slices"
	"strings"
)

func parseSelect(command []string, cfg *CliConfig) ([]string, error) {
	selectIndex := slices.Index(command, "select")
	if selectIndex == -1 {
		return []string{}, errors.New("no select statement found")
	}

	args := []string{}
	endingIndex := 0

	for i := selectIndex + 1; i < len(command); i++ {
		if slices.Contains(keywords, command[i]) {
			endingIndex = i - 1
			break
		}

		if i == len(command)-1 {
			if command[i][len(command[i])-1] == ',' {
				return []string{}, errors.New("unexpected ',' in select statement")
			}

			args = append(args, command[i])
			continue
		}

		if command[i][len(command[i])-1] != ',' && !slices.Contains(keywords, command[i+1]) {
			return []string{}, errors.New("missing a ',' in select statement")
		}

		args = append(args, strings.TrimRight(command[i][:len(command[i])], ","))
	}

	if endingIndex == 0 {
		endingIndex = len(command) - 1
	}

	cfg.Select = args

	return slices.Concat(command[:selectIndex], command[endingIndex+1:]), nil
}
