package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
)

const (
	trade           = 0
	topOfBook       = 30
	level2Orderbook = 8
)

type message struct {
	Action string   `json:"action"`
	APIKey string   `json:"api_key"`
	Subs   []string `json:"subs"`
}

type subscription struct {
	channel    int
	exchange   string
	marketFrom string
	marketTo   string
}

func generateFormattedSubscription(sub subscription) string {

	subMessage := strconv.Itoa(sub.channel)
	subMessage += `~`
	subMessage += sub.exchange
	if len(sub.marketFrom) != 0 && len(sub.marketTo) != 0 {
		// optionally add market if supplied
		subMessage += `~`
		subMessage += sub.marketFrom
		subMessage += `~`
		subMessage += sub.marketTo
	}
	return subMessage
}

func subscribe(subs []subscription, apiKey string) (string, error) {

	subscribeMessage := message{
		Action: "SubAdd",
		APIKey: apiKey,
	}

	for _, sub := range subs {
		subscribeMessage.Subs = append(subscribeMessage.Subs, generateFormattedSubscription(sub))
	}

	b, err := json.Marshal(subscribeMessage)
	if err != nil {
		return "", fmt.Errorf("marshalling subscribeMessage, %s", err)
	}

	return string(b), nil
}

func unsubscribe(subs []subscription, apiKey string) (string, error) {

	unsubscribeMessage := message{
		Action: "SubRemove",
		APIKey: apiKey,
	}

	for _, sub := range subs {
		unsubscribeMessage.Subs = append(unsubscribeMessage.Subs, generateFormattedSubscription(sub))
	}

	b, err := json.Marshal(unsubscribeMessage)
	if err != nil {
		return "", fmt.Errorf("marshalling unsubscribeMessage, %s", err)
	}

	return string(b), nil
}

func handleMessage(msg string) {
	log.Println("Received ", msg)
}

func main() {
	const url = "wss://streaming.cryptocompare.com"
	// Subscribe to trade, level 1 and level 2 channels for a single instrument
	subs := []subscription{
		subscription{
			channel:    trade,
			exchange:   "Coinbase",
			marketFrom: "BTC",
			marketTo:   "USD",
		},
		subscription{
			channel:    topOfBook,
			exchange:   "Coinbase",
			marketFrom: "BTC",
			marketTo:   "USD",
		},
		subscription{
			channel:    level2Orderbook,
			exchange:   "Coinbase",
			marketFrom: "BTC",
			marketTo:   "USD",
		},
	}
	subsMessage, err := subscribe(subs, "YOUR-API-KEY")
	if err != nil {
		log.Fatal(err)
	}

	interrupt := make(chan os.Signal, 1)
	done := make(chan struct{})

	conns, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conns.Close()

	// Send subscribe message on connection
	conns.WriteMessage(websocket.BinaryMessage, []byte(subsMessage))

	go func() {
		for {
			_, msg, errRead := conns.ReadMessage()
			if errRead != nil {
				log.Println("ReadMessage error ", errRead)
				_, ok := <-done
				if ok {
					close(done)
				}
			} else {
				handleMessage(string(msg))
			}
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			err := conns.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure,
					""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
		case <-done:
			return
		}
	}
}
