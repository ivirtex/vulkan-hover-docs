package cmd

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [URL]",
	Short: "Generate Markdown documentation from provided URL link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		if err := Generate(url); err != nil {
			cmd.Println("Error generating documentation:", err)
		} else {
			cmd.Println("Documentation generated successfully.")
		}
	},
}

func Generate(url string) error {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println("Found link:", e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	return nil
}
