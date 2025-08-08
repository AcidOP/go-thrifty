package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AcidOP/go-thrifty/pkg/notification"
	"github.com/AcidOP/go-thrifty/pkg/scraper"
	"github.com/joho/godotenv"
)

func main() {
	now := time.Now() // Use to account for the time of execution

	godotenv.Load()

	accountSID := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	number := os.Getenv("RECEIVER_PHONE_NUMBER")
	if len(accountSID) == 0 || len(authToken) == 0 {
		fmt.Println("Please set ACCOUNT_SID and AUTH_TOKEN in your environment variables.")
		return
	}

	notifier := notification.New(authToken, accountSID)

	products := []*scraper.Product{
		scraper.New("BOULT MUSTANG TORQ", "B0DN1Q4NSJ", 1600),
	}

	for _, prod := range products {
		currPrice, err := prod.Scrape()
		if err != nil {
			fmt.Println("Error scraping product:", err)
			continue
		}

		if currPrice <= prod.Price {
			notifier.Alert(number, prod)
		}
	}

	fmt.Println("Scraping completed in", time.Since(now))
}
