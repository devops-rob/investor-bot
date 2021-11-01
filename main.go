package main

import (
	"net/http"
	S "twitch-investo-bot/webserver"
)

func main()  {
	srv := S.NewServer()
	http.ListenAndServe(":80", srv)
}