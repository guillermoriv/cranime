package extractor

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type AnimeDescriptor struct {
	Title       string
	Description string
	Score       string
}

// Extractor for extracting the meta-data information of the supported
// URL links by the extractor.
func ExtractorMyAnimeList(query string) (*AnimeDescriptor, error) {
	c := colly.NewCollector()

	// we allocate a new descriptor to return with
	descriptor := new(AnimeDescriptor)

	c.OnHTML("div.score-label", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		descriptor.Score = e.Text
	})

	c.OnHTML("h1.title-name", func(e *colly.HTMLElement) {
		descriptor.Title = e.Text
	})

	c.OnHTML("p[itemprop=\"description\"]", func(e *colly.HTMLElement) {
		descriptor.Description = e.Text
	})

	err := c.Visit("https://myanimelist.net/" + query)

	// we catch the error in the collector if any
	if err != nil {
		return nil, err
	}

	return descriptor, nil
}
