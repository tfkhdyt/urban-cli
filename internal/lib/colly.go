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
}

func Scrape(word string, max int) ([]Result, error) {
	c := colly.NewCollector()
	c.MaxDepth = (max-1)/7 + 1

	var globalErr error
	c.OnError(func(r *colly.Response, err error) {
		globalErr = err
	})

	matchedWords := make([]string, 0)
	meanings := make([]string, 0)
	examples := make([]string, 0)
	contributors := make([]string, 0)

	c.OnHTML("a.word.text-denim.font-bold", func(h *colly.HTMLElement) {
		matchedWords = append(matchedWords, h.Text)
	})

	c.OnHTML("div.break-words.meaning.mb-4", func(h *colly.HTMLElement) {
		meanings = append(meanings, h.Text)
	})

	c.OnHTML("div.break-words.example.italic.mb-4", func(h *colly.HTMLElement) {
		examples = append(examples, h.Text)
	})

	c.OnHTML("div.contributor.font-bold", func(h *colly.HTMLElement) {
		contributors = append(contributors, h.Text)
	})

	c.OnHTML("a[aria-label='Next page']", func(h *colly.HTMLElement) {
		h.Request.Visit(h.Attr("href"))
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

	results := make([]Result, 0)
	for i := 0; i < len(meanings); i++ {
		results = append(results, Result{
			ID:          i + 1,
			Word:        strings.TrimSpace(matchedWords[i]),
			Meaning:     strings.TrimSpace(meanings[i]),
			Example:     strings.TrimSpace(examples[i]),
			Contributor: formatUsenameAndDate(contributors[i]),
		})
	}

	return results, nil
}
