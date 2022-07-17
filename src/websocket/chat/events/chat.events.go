package events

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func Connection(s socketio.Conn) error {
	s.SetContext("")
	fmt.Println("connected:", s.ID())
	return nil
}

func EventError(s socketio.Conn, e error) {
	fmt.Println("meet error:", e)
}

func Notice(s socketio.Conn, msg string) {
	fmt.Println("notice:", msg)
	s.Emit("reply", "have "+msg)
}

func Message(s socketio.Conn, msg string) string {
	fmt.Println("message:", msg)
	s.SetContext(msg)
	return "recv " + msg
}

func Bye(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("bye", last)
	s.Close()
	return last
}

func Disconnect(s socketio.Conn, reason string) {
	fmt.Println("closed", reason)
}
