package main

// Package is called aw
import (
	//"encoding/json"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// A Script Filter is required to return an items array of zero or more items.
// Each item describes a result row displayed in Alfred.
// https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	// The title displayed in the result row. There are no options for this element
	// and it is essential that this element is populated.
	Title string `json:"title"`
	// The subtitle displayed in the result row. This element is optional.
	SubTitle string `json:"subtitle"`
	// The argument which is passed through the workflow to the connected output action.
	Arg string `json:"arg"`
	// The icon displayed in the result row. Workflows are run from their workflow folder,
	// so you can reference icons stored in your workflow relatively.
	Icon string `json:"icon"`
}

func pageTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		return "", errors.Wrap(err, "Invalid HTTP status code")
	} else if err != nil {
		return "", errors.Wrap(err, "Invalid HTTP response")
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to load HTML document")
	}

	title := doc.Find("title").Text()

	return title, nil
}

func markdownLink(title, url string) string {
	return fmt.Sprintf("[%s](%s)", title, url)
}

func main() {
	var url *string = flag.String("url", "", "Generate GitHub flavored markdown link from a given URL")
	flag.Parse()

	title, err := pageTitle(*url)
	if err != nil {
		log.Fatal(err)
	}

	link := markdownLink(title, *url)

	item := Item{
		Title:    "Craft the specified URL as a markdown link",
		SubTitle: "Hit the 'Enter' key to copy this result to clipboard",
		Arg:      link,
		Icon:     "icon.png",
	}
	items := Items{Items: []Item{item}}

	jsonBytes, err := json.Marshal(items)
	if err != nil {
		log.Fatalf("JSON Marshal error: %s", err)
	}

	fmt.Println(string(jsonBytes))
}
