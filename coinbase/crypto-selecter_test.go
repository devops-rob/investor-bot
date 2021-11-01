package coinbase

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"twitch-investo-bot/mocks"
)

func init()  {
	Client = &mocks.MockClient{}
}

func TestOnlineFilter(t *testing.T) {

	products := []Product{
		Product{
			Id: "BAT-USDC",
			Status: "online",
		},
		Product{
			Id: "BAT-USDC",
			Status: "offline",

		},
	}

	result := OnlineFilter(products)

	assert.Len(t, result, 1)

}
