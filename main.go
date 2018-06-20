package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Define Websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func dbHandler(res http.ResponseWriter, req *http.Request) {
	var db_name string = req.URL.Path[1:]
	if len(db_name) > 0 {
		fmt.Fprintf(res, "Connecting to db %s...\n", db_name)
	} else {
		fmt.Fprintf(res, "No DB name given")
	}
}

func main() {
	http.HandleFunc("/", dbHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
