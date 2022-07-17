package chat

import (
	"websocket/src/websocket/chat/events"

	socketio "github.com/googollee/go-socket.io"
)

func WebSocketChat() *socketio.Server {
	server := socketio.NewServer(nil)
	ChatEvents(server)
	return server
}

func ChatEvents(socket *socketio.Server) {
	path := "/chat"

	socket.OnConnect("/", events.Connection)
	socket.OnError("/", events.EventError)
	socket.OnEvent("/", "notice", events.Notice)
	socket.OnEvent("/", "bye", events.Bye)
	socket.OnDisconnect("/", events.Disconnect)
	socket.OnEvent(path, "msg", events.Message)

}
