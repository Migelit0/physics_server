package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/migelit0/physics_server/WebSocket/server/config"
	core "github.com/migelit0/physics_server/core/structures"
	"log"
	"net"
	"os"
	"strconv"
)

var port int

func initWorld() core.World {
	var emptyBodies []core.Body
	w := core.World{Width: config.WIDTH, Height: config.HEIGHT, Bodies: emptyBodies, G: &config.G}
	return w
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	portStr, isExist := os.LookupEnv("PORT")
	if isExist {
		portEnv, err := strconv.Atoi(portStr)
		if err != nil {
			log.Println(err)
		}
		port = portEnv
	}

}

func main() {
	server := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	// serve conn somehow
}
