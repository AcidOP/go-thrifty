package scraper

import (
	"fmt"

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

	c := newAmazonCollector()

	c.OnHTML(".priceToPay .a-price-whole", func(e *colly.HTMLElement) {
		fmt.Println("💰 Price Found:", e.Text)
	})

	if err := c.Visit(url); err != nil {
		return fmt.Errorf("failed to visit %s: %w", url, err)
	}

	return nil
}
