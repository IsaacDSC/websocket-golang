// package main

// import (
// 	f "github.com/ambelovsky/gosf"
// )

// func init() {
// 	// Listen on an endpoint
// 	f.Listen("echo", func(client *f.Client, request *f.Request) *f.Message {
// 		f.)
// 		return f.NewSuccessMessage(request.Message.Text)
// 	})
// }

// func main() {
// 	// Start the server using a basic configuration
// 	f.Startup(map[string]interface{}{
// 		"port": 9999})
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	socketio "github.com/googollee/go-socket.io"
// )

// func main() {
// 	server := socketio.NewServer(nil)

// 	server.OnConnect("/", func(s socketio.Conn) error {
// 		s.SetContext("")
// 		fmt.Println("connected:", s.ID())
// 		return nil
// 	})

// 	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
// 		fmt.Println("notice:", msg)
// 		s.Emit("reply", "have "+msg)
// 	})

// 	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
// 		fmt.Println("message:", msg)
// 		s.SetContext(msg)
// 		return "recv " + msg
// 	})

// 	server.OnEvent("/", "bye", func(s socketio.Conn) string {
// 		last := s.Context().(string)
// 		s.Emit("bye", last)
// 		s.Close()
// 		return last
// 	})

// 	server.OnError("/", func(s socketio.Conn, e error) {
// 		fmt.Println("meet error:", e)
// 	})

// 	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 		fmt.Println("closed", reason)
// 	})

// 	go server.Serve()
// 	defer server.Close()

// 	http.Handle("/socket.io/", server)
// 	http.Handle("/", http.FileServer(http.Dir("./public")))
// 	log.Println("Serving at localhost:8000...")
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

package main

import (
	"websocket/src/server"
)

func main() {
	server.Server(8000)
}
