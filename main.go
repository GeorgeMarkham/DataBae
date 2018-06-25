package databae

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Define message type
type message struct {
	Method string
	Data   Record
}

//Define Websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func dbHandler(res http.ResponseWriter, req *http.Request) {
	//Grab Database name from the url
	var dbName = req.URL.Path[1:]

	//Get connection from the websocket uprader
	conn, err := upgrader.Upgrade(res, req, nil)

	//If there's an error print it to the console
	if err != nil {
		log.Println(err)
		return
	}
	//Check that a database name has been entered
	if len(dbName) > 0 {
		for {

			//Read in the message
			//messageType, msgBytes, err := conn.ReadMessage() //messageType = the type of message, string, close etc... ~ msgBytes = data stored as uint8[] ~ err = any errors

			msg := message{}

			err := conn.ReadJSON(&msg)

			//If there are any errors print them:
			if err != nil {
				log.Println(err)
			}

			//If there are no errors print the mesage and the client address to the console
			if err == nil {

			}
		}

	} else {
		//Create the error message as a byte array
		errMsg := []byte("No db specified")

		//Write the error message to the client & check for errors
		if err := conn.WriteMessage(1, errMsg); err != nil {
			//If there is an error then log it to the console
			log.Println(err)
		}
		//Close the connection because the client has not specified a database endpoint
		conn.Close()
	}
}

func main() {

	//Run the http handler concurrently using goroutine
	go http.HandleFunc("/", dbHandler)

	//Fatal is equivalent to Print() followed by a call to os.Exit(1).
	log.Fatal(http.ListenAndServe(":8080", nil))

}
