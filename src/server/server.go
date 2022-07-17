package server

import (
	"fmt"
	"log"
	"net/http"
	"websocket/src/websocket/events/chat"
)

func Server(port int) {
	server := chat.WebSocketChat()

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/healtcheck", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}))

	fmt.Println("Serving at localhost:" + fmt.Sprintf("%d", port))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
