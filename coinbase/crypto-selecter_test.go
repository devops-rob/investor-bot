package coinbase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"twitch-investo-bot/mocks"
)

func init()  {
	Client = &mocks.MockClient{}
}

func TestProductBookSuccess(t *testing.T) {
	jsonResponse := `[{"id":"BAT-USDC","base_currency":"BAT","quote_currency":"USDC","base_min_size":"1","base_max_size":"300000","quote_increment":"0.000001","base_increment":"0.000001","display_name":"BAT/USDC","min_market_funds":"1","max_market_funds":"100000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":false,"status":"online","status_message":"","auction_mode":false},{"id":"BTC-EUR","base_currency":"BTC","quote_currency":"EUR","base_min_size":"0.001","base_max_size":"10000","quote_increment":"0.01","base_increment":"0.00000001","display_name":"BTC/EUR","min_market_funds":"10","max_market_funds":"600000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":false,"status":"online","status_message":"","auction_mode":false},{"id":"ETH-BTC","base_currency":"ETH","quote_currency":"BTC","base_min_size":"0.01","base_max_size":"1000000","quote_increment":"0.00001","base_increment":"0.00000001","display_name":"ETH/BTC","min_market_funds":"0.001","max_market_funds":"80","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":false,"status":"online","status_message":"","auction_mode":false},{"id":"BTC-USD","base_currency":"BTC","quote_currency":"USD","base_min_size":"0.001","base_max_size":"10000","quote_increment":"0.01","base_increment":"0.00000001","display_name":"BTC/USD","min_market_funds":"10","max_market_funds":"1000000","margin_enabled":true,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":false,"status":"online","status_message":"","auction_mode":false},{"id":"BTC-GBP","base_currency":"BTC","quote_currency":"GBP","base_min_size":"0.001","base_max_size":"10000","quote_increment":"0.01","base_increment":"0.00000001","display_name":"BTC/GBP","min_market_funds":"10","max_market_funds":"200000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":false,"status":"online","status_message":"","auction_mode":false},{"id":"LINK-USDC","base_currency":"LINK","quote_currency":"USDC","base_min_size":"1","base_max_size":"800000","quote_increment":"0.000001","base_increment":"1","display_name":"LINK/USDC","min_market_funds":"10","max_market_funds":"100000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":false,"status":"online","status_message":"","auction_mode":false},{"id":"LINK-USD","base_currency":"LINK","quote_currency":"USD","base_min_size":"1","base_max_size":"800000","quote_increment":"0.000001","base_increment":"1","display_name":"LINK/USD","min_market_funds":"10","max_market_funds":"100000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":false,"status":"online","status_message":"","auction_mode":false},{"id":"XRP-GBP","base_currency":"XRP","quote_currency":"GBP","base_min_size":"1","base_max_size":"500000","quote_increment":"0.0001","base_increment":"0.000001","display_name":"XRP/GBP","min_market_funds":"10","max_market_funds":"100000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":true,"status":"delisted","status_message":"","auction_mode":false},{"id":"XRP-EUR","base_currency":"XRP","quote_currency":"EUR","base_min_size":"1","base_max_size":"500000","quote_increment":"0.0001","base_increment":"1","display_name":"XRP/EUR","min_market_funds":"10","max_market_funds":"100000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":true,"status":"delisted","status_message":"","auction_mode":false},{"id":"XRP-USD","base_currency":"XRP","quote_currency":"USD","base_min_size":"1","base_max_size":"500000","quote_increment":"0.0001","base_increment":"1","display_name":"XRP/USD","min_market_funds":"10","max_market_funds":"100000","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":true,"status":"delisted","status_message":"","auction_mode":false},{"id":"XRP-BTC","base_currency":"XRP","quote_currency":"BTC","base_min_size":"1","base_max_size":"500000","quote_increment":"0.00000001","base_increment":"1","display_name":"XRP/BTC","min_market_funds":"0.001","max_market_funds":"30","margin_enabled":false,"fx_stablecoin":false,"max_slippage_percentage":"0.10000000","post_only":false,"limit_only":false,"cancel_only":false,"trading_disabled":true,"status":"delisted","status_message":"","auction_mode":false}]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	resp, _ := ProductBook()

	if len(resp) == 0 {
		t.Error("TestProductBookSuccess failed, response was empty.")
		return
	}

}

func TestProductBookFail(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
			Body:       nil,
		}, errors.New("mock error")
	}

	_, err := ProductBook()
	if err == nil {
		t.Error("TestProductBookFail failed.")
		return
	}
}

func TestOnlineFilter(t *testing.T) {
	coins := `
[
  {
    "id": "BAT-USDC",
    "base_currency": "BAT",
    "quote_currency": "USDC",
    "base_min_size": "1",
    "base_max_size": "300000",
    "quote_increment": "0.000001",
    "base_increment": "0.000001",
    "display_name": "BAT/USDC",
    "min_market_funds": "1",
    "max_market_funds": "100000",
    "margin_enabled": false,
    "fx_stablecoin": false,
    "max_slippage_percentage": "0.10000000",
    "post_only": false,
    "limit_only": false,
    "cancel_only": false,
    "trading_disabled": false,
    "status": "online",
    "status_message": "",
    "auction_mode": false
  }
]`

	cryptoAssets := make([]Product,0)
	err := json.Unmarshal([]byte(coins), &cryptoAssets)
	if err != nil {
		t.Error("TestOnlineFilter failed. Unable to unmarshal json response")
		return
	}

	for _, v := range cryptoAssets {
		if v.Status != "online" {
			t.Error("TestOnlineFilter failed. Status should be online")
			return
		}
	}
}

func TestCurrencyFilter(t *testing.T) {
	coins := `
[
  {
    "id": "BAT-USDC",
    "base_currency": "BAT",
    "quote_currency": "USDC",
    "base_min_size": "1",
    "base_max_size": "300000",
    "quote_increment": "0.000001",
    "base_increment": "0.000001",
    "display_name": "BAT/USDC",
    "min_market_funds": "1",
    "max_market_funds": "100000",
    "margin_enabled": false,
    "fx_stablecoin": false,
    "max_slippage_percentage": "0.10000000",
    "post_only": false,
    "limit_only": false,
    "cancel_only": false,
    "trading_disabled": false,
    "status": "online",
    "status_message": "",
    "auction_mode": false
  }
]`

	cryptoAssets := make([]Product,0)
	err := json.Unmarshal([]byte(coins), &cryptoAssets)
	if err != nil {
		t.Error("TestOnlineFilter failed. Unable to unmarshal json response")
		return
	}

	for _, v := range cryptoAssets {
		if !strings.Contains(v.QuoteCurrency, "USD") {
			t.Error("TestOnlineFilter failed. Base currency should contain USD")
			return
		}
	}
}
func TestCryptoPicker(t *testing.T) {

}