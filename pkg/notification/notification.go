package notification

import (
	"fmt"

	"github.com/AcidOP/go-thrifty/pkg/scraper"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Notifier struct {
	client *twilio.RestClient
}

func New(authToken, accountSID string) *Notifier {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})

	return &Notifier{client: client}
}

// Alert sends a WhatsApp message to the specified number when the product's price drops below the alert price.
// It constructs a message with the product's name and a link to its Amazon page.
func (n *Notifier) Alert(number string, prod *scraper.Product) error {
	message := fmt.Sprintf(
		"Price Drop 💰 for %s\nLink: https://www.amazon.in/dp/%s",
		prod.Name, prod.ASIN,
	)
	return n.send(number, message)
}

// send sends a WhatsApp message to the specified number using Twilio's API.
func (n *Notifier) send(number, message string) error {
	if len(number) == 0 || len(message) == 0 {
		return fmt.Errorf("phone number and message must not be empty")
	}

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:" + number)
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody(message)

	if _, err := n.client.Api.CreateMessage(params); err != nil {
		return fmt.Errorf("NOTIFIER FAILED: %w", err)
	}
	return nil
}
