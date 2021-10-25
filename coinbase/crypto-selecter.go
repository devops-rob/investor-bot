package coinbase

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Product struct {
	Id           		  string `json:"id"`
	BaseCurrency 		  string `json:"base_currency,omitempty"`
	BaseMinSize    		  string `json:"base_min_size,omitempty"`
	BaseMaxSize    		  string `json:"base_max_size,omitempty"`
	QuoteCurrency		  string `json:"quote_currency,omitempty"`
	QuoteIncrement 		  string `json:"quote_increment,omitempty"`
	BaseIncrement  		  string `json:"base_increment,omitempty"`
	DisplayName    		  string `json:"display_name,omitempty"`
	MinMarketFunds        string `json:"min_market_funds,omitempty"`
	MaxMarketFunds        string `json:"max_market_funds,omitempty"`
	MarginEnabled         bool   `json:"margin_enabled,omitempty"`
	FxStablecoin          bool   `json:"fx_stablecoin,omitempty"`
	MaxSlippagePercentage string `json:"max_slippage_percentage,omitempty"`
	PostOnly              bool   `json:"post_only,omitempty"`
	LimitOnly             bool   `json:"limit_only,omitempty"`
	CancelOnly            bool   `json:"cancel_only,omitempty"`
	TradingDisabled       bool   `json:"trading_disabled,omitempty"`
	Status                string `json:"status,omitempty"`
	StatusMessage         string `json:"status_message,omitempty"`
}

type Crypto struct {
	Collection []Product
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func ProductBook() ([]Product, error) {
	url := "https://api-public.sandbox.pro.coinbase.com/products"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := Client.Do(request)
	if err != nil {
		return nil, err
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	cryptoAssets := make([]Product,0)
	json.Unmarshal(body, &cryptoAssets)

	return cryptoAssets, nil

}

func OnlineFilter() []Product  {
	// Filter assets for online trading status
	var onlineAsset []Product

	products, _ := ProductBook()

	for _, v := range products {
		if v.Status == "online" {
			onlineAsset = append(onlineAsset, v)
		}
	}

	return onlineAsset

}

func CurrencyFilter() []Product {
	// Filter online assets for USD quote currency
	var usdAsset []Product

	for _, v := range OnlineFilter() {
		if strings.Contains(v.QuoteCurrency, "USD") {
			usdAsset = append(usdAsset, v)
		}
	}
	return usdAsset
}

func CryptoPicker() Product {

	// Randomly select asset to buy from filtered slice

	rand.Seed(time.Now().Unix())

	chosenAsset := rand.Intn(len(CurrencyFilter()))
	pick := CurrencyFilter()[chosenAsset]

	return pick
}
