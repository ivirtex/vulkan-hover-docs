package cmd

import (
	"fmt"
	"io"
	"net/http"
	"sync"

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

type Page struct {
	URL     string
	Content string
	Err     error
}

func Generate(url string) error {
	var wg sync.WaitGroup
	ch := make(chan Page)

	// Create map to store page name to page content
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Println("Found link:", link)

		wg.Add(1)
		go DownloadPage(url+link, ch, &wg)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Gather results into a map
	contents := make(map[string]string)

	for page := range ch {
		if page.Err != nil {
			fmt.Printf("Failed to download %s: %v\n", page.URL, page.Err)
			continue
		}

		contents[page.URL] = page.Content
		fmt.Printf("Downloaded %s (%d bytes)\n", page.URL, len(page.Content))
	}

	fmt.Println("\n--- Summary ---")
	for url, content := range contents {
		fmt.Printf("%s => %d bytes\n", url, len(content))
	}

	if len(contents) == 0 {
		fmt.Println("No pages were downloaded.")
		return fmt.Errorf("no pages downloaded")
	} else {
		fmt.Println("Total pages downloaded:", len(contents))
		return nil
	}
}

func DownloadPage(url string, ch chan<- Page, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		ch <- Page{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Page{URL: url, Err: err}
		return
	}

	ch <- Page{URL: url, Content: string(body)}
}
