package main

// Kill all websocket command
// kill -9 $(lsof -i:5000 -t) 2> /dev/null

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello World")
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("New Connection")

		// Handle Chat messages
		so.Join("chat")

		so.On("chat message", func(msg string) {
			log.Println("Message Received from Client: " + msg)
			so.BroadcastTo("chat", "chat message", msg)
		})
	})

	fs := http.FileServer(http.Dir("static"))

	http.Handle("/", fs)

	http.Handle("/socket.io", server)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
