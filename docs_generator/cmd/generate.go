package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/charmbracelet/log"
	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("path", "p", "docs", "Path to directory where to save the generated Markdown files")
	generateCmd.Flags().BoolP("verbose", "v", false, "Enable verbose logging")
}

var generateCmd = &cobra.Command{
	Use:   "generate [URL]",
	Short: "Generate Markdown documentation from provided URL link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool("verbose")
		path, _ := cmd.Flags().GetString("path")
		url := RemoveSlashes(args[0])

		Generate(url, path, verbose)
	},
}

type Page struct {
	URL   string
	Title string
}

var postprocessPattern = `\[[^\]]+\]\((\w+.\w+)\)`
var re = regexp.MustCompile(postprocessPattern)

func Generate(url string, savePath string, verbose bool) {
	if url == "" {
		log.Fatal("URL cannot be empty")
		return
	}

	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// Start timer
	start := time.Now()

	pages, err := ExtractPages(url)
	if err != nil {
		log.Fatalf("Error extracting pages: %v", err)
		return
	}

	totalProcessedPages, totalSize := DownloadAndProcessPages(pages, url, savePath)

	elapsed := time.Since(start)
	log.Infof("Documentation generation completed in %s", elapsed)

	log.Infof("Generated %d/%d pages", totalProcessedPages, len(pages))
	log.Infof("Total size: %f MB", float32(totalSize)/1024.0/1024.0)
}

func ExtractPages(url string) ([]Page, error) {
	var pages []Page

	// Create map to store page name to page content
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		sanitizedLink := RemoveSlashes(RemoveParameters(link))

		// Filter out links that are not relevant
		ignoredNames := []string{"", "#", ".", "..", "?"}
		for _, ignored := range ignoredNames {
			if sanitizedLink == ignored {
				log.Infof("Ignoring link: %s", link)

				return
			}
		}

		log.Debugf("Found link: %s", sanitizedLink)
		pages = append(pages, Page{
			URL:   url + "/" + sanitizedLink,
			Title: RemoveExtension(sanitizedLink),
		})
	})

	log.Infof("Analyzing URL: %s", url)

	c.Visit(url)
	c.Wait()

	if len(pages) == 0 {
		return nil, fmt.Errorf("no pages found at the provided URL: %s", url)
	}

	log.Infof("Found %d pages at %s", len(pages), url)

	return pages, nil
}

const MaxConcurrentDownloads = 30

// Downloads and converts provided pages to Markdown concurrently
// Returns a map of Page to Markdown content
func DownloadAndProcessPages(pages []Page, originURL string, save_path string) (totalProcessedPages, totalSize int) {
	sem := make(chan struct{}, MaxConcurrentDownloads)

	for _, page := range pages {
		sem <- struct{}{} // Acquire a token
		go func(p Page) {
			defer func() { <-sem }() // Release the token when done

			log.Infof("Processing page: %s", p.URL)

			log.Debugf("Downloading page: %s", p.URL)
			htmlContent, err := DownloadPage(p.URL)
			if err != nil {
				log.Errorf("Error downloading page %s: %v\n", p.URL, err)
				return
			}

			log.Debugf("Converting page %s to Markdown", p.URL)
			mdContent, err := ConvertPage(htmlContent)
			if err != nil {
				log.Errorf("Error converting page %s to Markdown: %v\n", p.URL, err)
				return
			}

			log.Debugf("Post-processing page %s", p.Title+".md")
			mdContent, err = PostProcessPage(originURL, mdContent)
			if err != nil {
				log.Errorf("Error post-processing page %s: %v\n", p.Title+".md", err)
				return
			}

			log.Debugf("Saving page %s to %s", p.Title+".md", save_path)
			if err := SavePage(p, mdContent, save_path); err != nil {
				log.Errorf("Error saving page %s: %v\n", p.Title+".md", err)
				return
			}

			totalProcessedPages++
			totalSize += len(mdContent)
		}(page)
	}

	// Wait for all goroutines to finish
	for range MaxConcurrentDownloads {
		sem <- struct{}{} // Ensure we release all tokens
	}

	return
}

// Downloads the HTML content of the page at the given URL
func DownloadPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download page: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download page: %s", resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read page content: %w", err)
	}

	return string(content), nil
}

// Converts HTML content to Markdown
func ConvertPage(content string) (string, error) {
	markdown, err := htmltomarkdown.ConvertString(content)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to Markdown: %w", err)
	}

	return markdown, nil
}

// Removes unnecessary content and fixes links in the Markdown content
func PostProcessPage(originURL string, content string) (string, error) {
	// Remove "Loading... please wait." from the content
	content = strings.ReplaceAll(content, "Loading… please wait.", "")

	// Prepend main URL to the HTML links like:
	content = re.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the link from the group
		referencedPage := re.FindStringSubmatch(match)[1]
		referencedPageWithURL := ""

		// Prepend the main URL to the link
		if !strings.HasPrefix(referencedPage, "http") {
			referencedPageWithURL = originURL + "/" + referencedPage
		}

		// Return the modified match
		return strings.Replace(match, referencedPage, referencedPageWithURL, 1)
	})

	return content, nil
}

// Saves the Markdown content to a file
func SavePage(page Page, content string, path string) error {
	if path == "" {
		path = "docs"
	}

	// Ensure the directory exists
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}

	// Create a filename based on the page title or URL
	filename := page.Title + ".md"
	filePath := fmt.Sprintf("%s/%s", path, filename)

	// Write the content to the file
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to save page %s: %w", filePath, err)
	}

	return nil
}

func RemoveSlashes(str string) string {
	// Remove any leading or trailing slashes
	if len(str) > 0 && str[0] == '/' {
		str = str[1:]
	}

	if len(str) > 0 && str[len(str)-1] == '/' {
		str = str[:len(str)-1]
	}

	return str
}

func RemoveParameters(str string) string {
	if questionMark := strings.Index(str, "?"); questionMark != -1 {
		return str[:questionMark]
	}

	return str
}

func RemoveExtension(str string) string {
	if dot := strings.LastIndex(str, "."); dot != -1 {
		return str[:dot]
	}

	return str
}
