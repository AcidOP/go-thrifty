package main

import (
	"fmt"

	"github.com/AcidOP/go-thrifty/pkg/scraper"
)

func main() {
	products := []*scraper.Product{
		scraper.New("Keyboard", "B0DN1Q4NSJ", 1599),
	}

	for _, p := range products {
		if err := p.Scrape(); err != nil {
			fmt.Println("Error scraping product:", err)
		}
	}
}
