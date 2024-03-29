package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/cndpost/go-socket.io/v1.0"
)

func main() {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {

		log.Println("on connection")

		so.Join("chat")

		so.On("chat message", func(msg string) {
			log.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("notice", func(msg string) {
		fmt.Println("notice:", msg)
		//	s.Emit("reply", "have "+msg)
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)

	//	fs := http.FileServer(http.Dir("static"))
	//	chat := http.FileServer(http.Dir("./asset"))
	//	http.Handle("/", fs)
	//	http.Handle("/chat", chat)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	//	http.HandleFunc("/chat", http.FileServer(http.Dir("./asset")))

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
