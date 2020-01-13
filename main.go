package main

// Package is called aw
import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

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
		return "", errors.Errorf("Invalid HTTP status code: %d", resp.StatusCode)
	} else if err != nil {
		return "", errors.Wrap(err, "Invalid HTTP response")
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to load HTML document")
	}

	title := doc.Find("head > title").Text()

	return title, nil
}

func markdownLink(title, url string) string {
	return fmt.Sprintf("[%s](%s)", title, url)
}

func validateURL(arg string) error {
	// If there are any major problems with the format of the URL, url.Parse() will
	// return an error.
	u, err := url.Parse(arg)
	if err != nil {
		return errors.Wrapf(err, "%s is not a valid URL", arg)
	} else if u.Scheme == "" || u.Host == "" {
		return errors.Errorf("%s must be an absolute URL", arg)
	} else if u.Scheme != "http" && u.Scheme != "https" {
		return errors.Errorf("%s must begin with http or https", arg)
	}
	return nil
}

func main() {
	var url *string = flag.String("url", "", "Generate GitHub flavored markdown link from a given URL")
	flag.Parse()

	err := validateURL(*url)
	if err != nil {
		fmt.Println(err)
	}

	title, err := pageTitle(*url)
	if err != nil {
		fmt.Println(err)
	}

	link := markdownLink(title, *url)

	fmt.Println(link)
}
