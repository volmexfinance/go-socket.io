package main

import (
	"encoding/json"
	"log"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

func main() {
	jwtToken := "<JWT_TOKEN>"
	uri := "wss://ws.volmex.finance"

	client, err := socketio.NewClient(uri, &engineio.Options{
		Transports: []transport.Transport{websocket.Default},
		JwtToken:   jwtToken,
	})
	if err != nil {
		panic(err)
	}

	client.OnEvent("indices-messages-stream-private", func(s socketio.Conn, msg json.RawMessage) {
		var indicesMessage struct {
			Symbol    string  `json:"symbol"`
			Price     float64 `json:"price"`
			Timestamp int64   `json:"timestamp"`
		}
		err := json.Unmarshal(msg, &indicesMessage)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
		} else {
			log.Printf("IndicesMessage: %v", indicesMessage)
		}
	})

	err = client.Connect()
	if err != nil {
		panic(err)
	}

	client.OnConnect(func(s socketio.Conn) error {
		log.Println("Connected to server")
		s.Emit("fetch-indices-messages-private")
		go func() {
			for {
				time.Sleep(5 * time.Second)
			}
		}()
		return nil
	})

	time.Sleep(1000 * time.Second)
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
