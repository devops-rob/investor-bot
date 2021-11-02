package main

import (
	"fmt"
	"log"
	"net/http"
	S "twitch-investo-bot/webserver"
)

func main()  {
	srv := S.NewServer()
	
	fmt.Printf("Starting server at: http://localhost:3000\n")
	if err := http.ListenAndServe(":3000", srv); err != nil {
		log.Fatal(err)
	}
}