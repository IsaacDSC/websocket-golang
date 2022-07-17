package server

import (
	"fmt"
	"log"
	"net/http"
	"websocket/src/routes"
	"websocket/src/websocket/chat"
)

func Server(port int) {
	server := chat.WebSocketChat()

	go server.Serve()
	defer server.Close()

	routes.SettingsRoutes(server)

	fmt.Println("Serving at localhost:" + fmt.Sprintf("%d", port))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
