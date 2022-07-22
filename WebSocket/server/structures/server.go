package structures

import (
	"fmt"
	"github.com/gorilla/websocket"
	core "github.com/migelit0/physics_server/core/structures"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

type Server struct {
	clients       map[*websocket.Conn]bool
	worlds        map[*websocket.Conn]core.World
	handleMessage func(message []byte) // хандлер новых сообщений,
	width, height uint16
	g float64
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	var newWorld core.World = core.World{
		Width:  server.width,
		Height: server.height,
		Bodies: []core.Body,
		G:      &server.g,
	}

	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	server.clients[connection] = true // Сохраняем соединение, используя его как ключ
	server.worlds[connection] = newWorld

	defer delete(server.clients, connection) // Удаляем соединение

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		go server.handleMessage(message)
	}
}

func (server *Server) WriteMessage(message []byte) {
	for conn := range server.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

func StartServer(handleMessage func(message []byte), port int) *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
		make(map[*websocket.Conn]core.World),
		handleMessage,
	}

	http.HandleFunc("/", server.echo)

	portStr := fmt.Sprintf(":%d", port)
	go http.ListenAndServe(portStr, nil) // Уводим http сервер в горутину

	return &server
}
