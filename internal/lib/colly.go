package lib

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Result struct {
	ID          int
	Word        string
	Meaning     string
	Example     string
	Contributor string
	Autolinks   []string
}

func Scrape(word string, max int) ([]Result, error) {
	c := colly.NewCollector()
	c.MaxDepth = (max-1)/7 + 1

	var globalErr error
	c.OnError(func(r *colly.Response, err error) {
		globalErr = err
	})

	results := make([]Result, 0)

	c.OnHTML("div.definition", func(h *colly.HTMLElement) {
		r := Result{
			ID:          len(results) + 1,
			Word:        h.ChildText("a.word"),
			Meaning:     h.ChildText("div.meaning"),
			Example:     h.ChildText("div.example"),
			Contributor: formatUsernameAndDate(h.ChildText("div.contributor")),
			Autolinks:   h.ChildTexts("a.autolink"),
		}
		results = append(results, r)
	})

	c.OnHTML("a[aria-label='Next page']", func(h *colly.HTMLElement) {
		_ = h.Request.Visit(h.Attr("href"))
	})

	// for i := 1; i <= (max-1)/7+1; i++ {
	url := fmt.Sprintf(
		"https://www.urbandictionary.com/define.php?term=%s",
		strings.ReplaceAll(word, " ", "+"),
	)

	if err := c.Visit(url); err != nil {
		return nil, fmt.Errorf("something went wrong: %w", err)
	}
	// }

	if globalErr != nil {
		return nil, fmt.Errorf("something went wrong: %w", globalErr)
	}

	return results, nil
}
