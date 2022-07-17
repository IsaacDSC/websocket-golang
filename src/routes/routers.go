package routes

import (
	"log"
	"net/http"
	"os"
	"websocket/src/controllers"
	"websocket/src/middlewares"

	socketio "github.com/googollee/go-socket.io"
)

func SettingsRoutes(server *socketio.Server) {
	http.Handle("/socket.io/", (server))
	routes()
}

func routes() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	http.Handle("/test", middlewares.Notify(logger)(http.HandlerFunc(controllers.HealthCheck)))
	http.Handle("/healtcheck", http.HandlerFunc(controllers.HealthCheck))
}
