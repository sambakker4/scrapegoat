/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var url string
var rootCmd = &cobra.Command{
	Use:   "scrapegoat",
	Short: "A webscraper that scrapes for dead links",
	Long:  `Scrapegoat is a web scraper that recursely searches for dead links and prints them to the console`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		deadLinks, err := ProcessURL(url)
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			return
		}

		fmt.Println()
		PrintDeadLinks(deadLinks)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "The url to searched for dead links")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
