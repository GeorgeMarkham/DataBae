package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Define message type
type message struct {
	Code    int
	Command string
	Options string
}

//Define Websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func dbHandler(res http.ResponseWriter, req *http.Request) {
	var dbName = req.URL.Path[1:]
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	if len(dbName) > 0 {
		for {
			messageType, msgBytes, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
			}
			log.Printf("Message from %s: %s", conn.RemoteAddr(), string(msgBytes))
			if err := conn.WriteMessage(messageType, msgBytes); err != nil {
				log.Println(err)
				return
			}
			if messageType == websocket.CloseMessage {
				log.Printf("Connection closed by client: %s", conn.RemoteAddr())
			}
		}

	} else {
		errMsg := []byte("No db specified")
		if err := conn.WriteMessage(1, errMsg); err != nil {
			log.Println(err)
		}
		conn.Close()
	}
}

func main() {
	go http.HandleFunc("/", dbHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
