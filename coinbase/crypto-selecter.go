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

func cryptoPicker() Product {
	// API call to get list of Crypto products
	resp, err := http.Get("https://api-public.sandbox.pro.coinbase.com/products")
	if err != nil {
		log.Fatalln(err)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Unmarshal json response into Product struct
	cryptoAssets := make([]Product,0)
	json.Unmarshal(body, &cryptoAssets)

	// Filter assets for online trading status
	var onlineAsset []Product

	for _, v := range cryptoAssets {
		if v.Status == "online" {
			onlineAsset = append(onlineAsset, v)
		}
	}

	// Filter online assets for USD quote currency
	var usdAsset []Product

	for _, v := range onlineAsset {
		if strings.Contains(v.QuoteCurrency, "USD") {
			usdAsset = append(usdAsset, v)
		}
	}

	// Randomly select asset to buy from filtered slice

	rand.Seed(time.Now().Unix())

	chosenAsset := rand.Intn(len(usdAsset))
	pick := usdAsset[chosenAsset]

	return pick
}
