package main

import (
	"net/http"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type config struct {
	url string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) < 2 {
		fmt.Println("You must provide a url")
		os.Exit(1)
	}

	url := os.Args[1]
	if !isValidURL(url) {
		fmt.Println("Invalid URL")
		os.Exit(1)
	}

	for {
		fmt.Printf("> ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

	}

}

func newCollector() *colly.Collector {
	collector := colly.NewCollector()

	collector.OnHTML("h1", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
	})

	collector.OnRequest(func(req *colly.Request) {
		fmt.Println("Visting:", req.URL)
	})

	return collector
}

func cleanInput(text string) []string {
	lowercaseText := strings.ToLower(text)
	return strings.Fields(lowercaseText)
}

func isValidURL(url string) bool {
	client := &http.Client{}

	resp, err := client.Head(url)
	if err != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}
