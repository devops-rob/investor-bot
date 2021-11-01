package coinbase

import (
	"fmt"
	"github.com/preichenberger/go-coinbasepro/v2"
)

func Invest()  {
	// Create a new client to auth with Coinbase Pro
	client := coinbasepro.NewClient()

	products, err := ProductBook()
	if err != nil {
		println(err.Error())
	}
	products = OnlineFilter(products)
	products = CurrencyFilter(products, "USD")
	product := CryptoPicker(products)

	// Place the order on Coinbase Pro
	order := coinbasepro.Order{
		Funds:     product.MinMarketFunds,
		Side:      "buy",
		ProductID: product.Id,
		Type:      "market",
	}

	savedOrder, err := client.CreateOrder(&order)
	if err != nil {
		println(err.Error())
	}

	// Print the Crypto asset that was purchased
	fmt.Println(product.Id)

	// Print the Order ID
	fmt.Printf("Order ID: %s", savedOrder.ID)
}
