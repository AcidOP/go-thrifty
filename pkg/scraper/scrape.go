package scraper

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Name  string
	ASIN  string
	Price int
}

func New(name, asin string, price int) *Product {
	return &Product{
		Name:  name,
		ASIN:  asin,
		Price: price,
	}
}

func (p *Product) Scrape() error {
	url := "https://www.amazon.in/dp/" + p.ASIN

	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.in", "amazon.in"),
	)

	c.SetRequestTimeout(30 * time.Second)

	// Amazon blocks requests that do noy have a User-Agent header set.
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")

		fmt.Printf("\n[Visiting]: %s\n\n", r.URL)
	})

	c.OnHTML(".priceToPay .a-price-whole", func(e *colly.HTMLElement) {
		fmt.Println("✅ Final Product Price:", e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("\n[Error] %s: %v\n", r.Request.URL, err)
	})

	if err := c.Visit(url); err != nil {
		return fmt.Errorf("failed to visit URL %s: %w", url, err)
	}

	return nil
}
