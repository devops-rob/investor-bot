package webserver

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"twitch-investo-bot/coinbase"
)

type Server struct {
	*mux.Router
	Config Config
}
type Config struct {
	Event        Event        `json:"event"`
	Subscription Subscription `json:"subscription"`
	Challenge    string       `json:"challenge,omitempty"`
}
type Event struct {
	UserId               string `json:"user_id,omitempty"`
	UserLogin            string `json:"user_login,omitempty"`
	UserName             string `json:"user_name,omitempty"`
	BroadcasterUserId    string `json:"broadcaster_user_id,omitempty"`
	BroadcasterUserLogin string `json:"broadcaster_user_login,omitempty"`
	BroadcasterUserName  string `json:"broadcaster_user_name,omitempty"`
}

type Subscription struct {
	Id        string `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	Type      string `json:"type,omitempty"`
	Version   string `json:"version,omitempty"`
	Cost      int    `json:"cost,omitempty"`
	Condition struct {
		BroadcasterUserId string `json:"broadcaster_user_id,omitempty"`
	} `json:"condition,omitempty"`
	Transport struct {
		Method   string `json:"method,omitempty"`
		Callback string `json:"callback,omitempty"`
	} `json:"transport,omitempty"`
}

type MyResponseWriter struct {
	http.ResponseWriter
	buf *bytes.Buffer
}

func NewServer() *Server  {
	s := &Server{
		Router: mux.NewRouter(),
		Config: Config{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/", s.rootCallback())
	s.HandleFunc("/twitch-callback", s.twitchCallback()).Methods("POST")
}

func (s *Server) rootCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\":\"Welcome to the Twitch Investo-bot!\"}"))
	}
}

func (s *Server) twitchCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Read request body and store in variable
		requestData,err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Unmarshal request body into struct
		var c Config
		err = json.Unmarshal(requestData, &c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// Prepare request body, message ID and Timestamp for HMAC verification
		requestDataByte := []byte(requestData)

		twitchSignature := r.Header.Get("Twitch-Eventsub-Message-Signature")
		msgId := r.Header.Get("Twitch-Eventsub-Message-Id")
		msgTimestamp := r.Header.Get("Twitch-Eventsub-Message-Timestamp")

		// Check if time stamp is older than 10 minutes
		currentTimeStamp := time.Now()
		t, err := time.Parse("2006-01-02T15:04:05.000000Z", msgTimestamp)
		if err != nil {
			fmt.Println(err)
		}
		difference := currentTimeStamp.Sub(t)
		if difference > time.Minute *10 {
			w.WriteHeader(http.StatusForbidden)
		}

		var secret = "123456789"
		key := []byte(secret)

		hmacMessage := fmt.Sprintf("%s%s%s", msgId, msgTimestamp, requestDataByte)

		// Perform HMAC on concatenated payload
		mac := hmac.New(sha256.New, key)
		mac.Write([]byte(hmacMessage))

		computedSignature := "sha256="+hex.EncodeToString(mac.Sum(nil))

		// Compare computed HMAC signature to Twitch header signature
		if computedSignature != twitchSignature {
			w.WriteHeader(http.StatusForbidden)
			forbidden := http.StatusForbidden
			log.Println(forbidden)
		}

		if r.Header.Get("Twitch-Eventsub-Message-Type") == "webhook_callback_verification" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(c.Challenge))
			return
		}

		if r.Header.Get("Twitch-Eventsub-Message-Type") == "notification" {

			log.Printf("%s notification received from %s. Message ID: %s",
				r.Header.Get("Twitch-Eventsub-Subscription-Type"),
				c.Event.UserId,
				r.Header.Get("Twitch-Eventsub-Message-Id"),
			)
			log.Println("Purchasing Crypto")
			coinbase.Invest()
		}
	}
}

