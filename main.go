package main

import (
	"net/http"
	S "twitch-investo-bot/webserver"
)

func main()  {
	//coinbase.Invest()
	srv := S.NewServer()
	http.ListenAndServe(":80", srv)
}