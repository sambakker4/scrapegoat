package cmd

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
)

const minCellWidth = 0
const tabWidth = 2
const padding = 4
const padchar = ' '
const flags = 0

func ProcessURL(url string) ([][]string, error) {
	if url == "" {
		return [][]string{}, errors.New("no url provided, please provide a url with the --url flage")
	}
	return [][]string{
		{"http://google.com", "hey there"},
		{"http://google.com", "hey"},
	}, nil
}

func PrintLinks(links [][]string) {
	writer := tabwriter.NewWriter(os.Stdout, minCellWidth, tabWidth, padding, padchar, flags)
	_, err := writer.Write([]byte("Page\tLink\n"))
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	for _, row := range links {
		_, err = writer.Write(fmt.Appendf([]byte{}, "%s\t%s\n", row[0], row[1]))
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}
	}

	writer.Flush()
}
