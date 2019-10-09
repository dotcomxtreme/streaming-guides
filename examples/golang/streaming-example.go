package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gorilla/websocket"
)

const (
	Trade           = 0
	TopOfBook       = 30
	Level2Orderbook = 8
)

type Message struct {
	action  string
	api_key string
	subs    []string
}

type Subscription struct {
	channel     int
	exchange    string
	market_from string
	market_to   string
}

func GenerateFormattedSubscription(sub Subscription) string {

	subMessage := `"`
	subMessage += strconv.Itoa(sub.channel)
	subMessage += `~`
	subMessage += sub.exchange
	if len(sub.market_from) != 0 && len(sub.market_to) != 0 {
		// optionally add market if supplied
		subMessage += `~`
		subMessage += sub.market_from
		subMessage += `~`
		subMessage += sub.market_to
	}
	subMessage += `"`
	return subMessage
}

func Subscribe(subs []Subscription, apiKey string) string {
	subscribeMessage := `{"action":"SubAdd", "api_key":"`
	subscribeMessage += apiKey
	subscribeMessage += `", "subs":[`

	for index, sub := range subs {
		if index != 0 {
			subscribeMessage += `,`
		}
		subscribeMessage += GenerateFormattedSubscription(sub)
	}

	subscribeMessage += `]}`
	return subscribeMessage
}

func Unsubscribe(subs []Subscription, apiKey string) {
	subscribeMessage := `{"action":"SubRemove", "api_key":"`
	subscribeMessage += apiKey
	subscribeMessage += `", "subs":[`

	for index, sub := range subs {
		if index != 0 {
			subscribeMessage += `,`
		}
		subscribeMessage += GenerateFormattedSubscription(sub)
	}

	subscribeMessage += `]}`
}

func HandleMessage(msg string) {
	log.Println("Received ", msg)
}

func main() {
	const url = "wss://streaming.cryptocompare.com"
	// Subscribe to trade, level 1 and level 2 channels for a single instrument
	subs := []Subscription{
		Subscription{
			channel:     Trade,
			exchange:    "Coinbase",
			market_from: "BTC",
			market_to:   "USD",
		},
		Subscription{
			channel:     TopOfBook,
			exchange:    "Coinbase",
			market_from: "BTC",
			market_to:   "USD",
		},
		Subscription{
			channel:     Level2Orderbook,
			exchange:    "Coinbase",
			market_from: "BTC",
			market_to:   "USD",
		},
	}
	subsMessage := Subscribe(subs, "YOUR-API-KEY")

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
				HandleMessage(string(msg))
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
