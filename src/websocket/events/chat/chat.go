package chat

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func WebSocketChat() *socketio.Server {
	server := socketio.NewServer(nil)
	ChatEvents(server)
	return server
}

func ChatEvents(socket *socketio.Server) {
	path := "/chat"

	socket.OnConnect("/", connection)
	socket.OnError("/", eventError)
	socket.OnEvent("/", "notice", notice)
	socket.OnEvent("/", "bye", bye)
	socket.OnDisconnect("/", disconnect)

	socket.OnEvent(path, "msg", message)

}

func connection(s socketio.Conn) error {
	s.SetContext("")
	fmt.Println("connected:", s.ID())
	return nil
}

func eventError(s socketio.Conn, e error) {
	fmt.Println("meet error:", e)
}

func notice(s socketio.Conn, msg string) {
	fmt.Println("notice:", msg)
	s.Emit("reply", "have "+msg)
}

func message(s socketio.Conn, msg string) string {
	fmt.Println("message:", msg)
	s.SetContext(msg)
	return "recv " + msg
}

func bye(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("bye", last)
	s.Close()
	return last
}

func disconnect(s socketio.Conn, reason string) {
	fmt.Println("closed", reason)
}
