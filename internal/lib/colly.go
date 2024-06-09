package lib

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
		meaning, err := parseContent(h, "div.meaning")
		if err != nil {
			globalErr = err
			return
		}

		example, err := parseContent(h, "div.example")
		if err != nil {
			globalErr = err
			return
		}

		r := Result{
			ID:          len(results) + 1,
			Word:        h.ChildText("a.word"),
			Meaning:     meaning,
			Example:     example,
			Contributor: formatUsernameAndDate(h.ChildText("div.contributor")),
			Autolinks:   h.ChildTexts("a.autolink"),
		}
		results = append(results, r)
	})

	c.OnHTML("a[aria-label='Next page']", func(h *colly.HTMLElement) {
		_ = h.Request.Visit(h.Attr("href"))
	})

	url := fmt.Sprintf(
		"https://www.urbandictionary.com/define.php?term=%s",
		strings.ReplaceAll(word, " ", "+"),
	)

	if err := c.Visit(url); err != nil {
		return nil, err
	}

	if globalErr != nil {
		return nil, fmt.Errorf("something went wrong: %w", globalErr)
	}

	return results, nil
}

func parseContent(h *colly.HTMLElement, selector string) (string, error) {
	str, err := h.DOM.Find(selector).Html()
	if err != nil {
		return "", fmt.Errorf("failed to get selector: %w", err)
	}
	raw := strings.ReplaceAll(str, "<br/><br/>", "\n")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(raw))
	if err != nil {
		return "", fmt.Errorf("failed to parse document: %w", err)
	}

	return doc.Text(), nil
}
