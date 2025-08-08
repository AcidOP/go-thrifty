package scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var HEADERS = map[string]string{
	"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
	"Accept-Language": "en-US,en;q=0.9",
}

func newAmazonCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.in", "amazon.in"),
	)
	c.SetRequestTimeout(30 * time.Second)

	c.OnRequest(func(r *colly.Request) {
		for key, value := range HEADERS {
			r.Headers.Set(key, value)
		}

		fmt.Printf("🔍 Visiting: %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("❌ Error scraping %s: %v\n", r.Request.URL, err)
	})

	return c
}
