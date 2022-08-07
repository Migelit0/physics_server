package structures

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/migelit0/physics_server/WebSocket/server/util"
	core "github.com/migelit0/physics_server/core/structures"
	"log"
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
	g             float64
	Port          string
}

func (server *Server) handle(w http.ResponseWriter, r *http.Request) {
	var newWorld = util.GenerateInitWorld(3)

	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	server.clients[connection] = true    // Сохраняем соединение, используя его как ключ
	server.worlds[connection] = newWorld // Создаем мир для нового состояния

	defer delete(server.clients, connection) // Удаляем соединение
	defer delete(server.worlds, connection)  // Удаляем мир

	log.Println("New conn")
	for {
		mt, _, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		// go server.handleMessage(message)
		if mt == websocket.TextMessage && true {
			// нас просто пингуют - отдаем состояние мира

			// FIXME: адекватно сделать без лишней переменной
			world := server.worlds[connection]
			world.DoOneIter()
			server.worlds[connection] = world
			log.Println("World len(body): ", len(world.Bodies))

			msg := generateResponseText(&world)

			log.Println("sending: ", string(msg))
			err := connection.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}
		}
	}
}

func StartServer(handleMessage func(message []byte), port int, width, height uint16, g float64) *Server {
	portStr := fmt.Sprintf(":%d", port)

	server := Server{
		make(map[*websocket.Conn]bool),
		make(map[*websocket.Conn]core.World),
		handleMessage,
		width, height, g, portStr,
	}

	http.HandleFunc("/v1", server.handle)
	http.HandleFunc("/test", server.echo)

	go http.ListenAndServe(portStr, nil) // Уводим http сервер в горутину

	return &server
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	log.Println("In echo")
	connection, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Upgrade: ", err)
		return
	}

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			log.Println("read", err)
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		// go server.handleMessage(message)

		log.Println("send: ", string(message))
		err = connection.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}

	}
}
