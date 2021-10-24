package coinbase

import (
	"fmt"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func Invest()  {
	// Create a new client to auth with Coinbase Pro
	client := coinbasepro.NewClient()

	// Place the order on Coinbase Pro
	order := coinbasepro.Order{
		Funds:     CryptoPicker().MinMarketFunds,
		Side:      "buy",
		ProductID: CryptoPicker().Id,
		Type:      "market",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		println(err.Error())
	}

	// Print the Crypto asset that was purchased
	fmt.Println(CryptoPicker().Id)

	// Print the Order ID
	fmt.Printf("Order ID: %s", savedOrder.ID)
}
