package scraper

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

// Product represents an Amazon product with its name, ASIN, and the price to alert on.
// ASIN is the unique identifier for the product on Amazon.
type Product struct {
	Name  string
	ASIN  string
	Price int // Will be alerted if the price drops below this value
}

func New(name, asin string, price int) *Product {
	return &Product{Name: name, ASIN: asin, Price: price}
}

// Scrape crunches the latest price found on the product's Amazon page.
// In case of an error, the price is set to 0.
func (p *Product) Scrape() (int, error) {
	url := fmt.Sprintf("https://www.amazon.in/dp/%s", p.ASIN)
	c := newAmazonCollector()
	foundPrice := ""

	c.OnHTML(".priceToPay .a-price-whole", func(e *colly.HTMLElement) {
		foundPrice = e.Text
	})

	if err := c.Visit(url); err != nil {
		return 0, fmt.Errorf("failed to visit %s: %w", url, err)
	}

	if len(foundPrice) == 0 {
		return 0, fmt.Errorf("no price found for %s", p.Name)
	}

	return convertPrice(foundPrice)
}

// convertPrice converts a price string (e.g., "1,699") to an integer (e.g., 1699).
func convertPrice(price string) (int, error) {
	buf := bytes.Buffer{}

	for _, digit := range price {
		if digit >= '0' && digit <= '9' {
			buf.WriteRune(digit)
		}
	}

	if buf.Len() == 0 {
		return 0, fmt.Errorf("no digits found in price string")
	}

	return strconv.Atoi(buf.String())
}
