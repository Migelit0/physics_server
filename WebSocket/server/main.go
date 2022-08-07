package main

import (
	"github.com/joho/godotenv"
	"github.com/migelit0/physics_server/WebSocket/server/config"
	"github.com/migelit0/physics_server/WebSocket/server/structures"
	"log"
	"os"
	"strconv"
)

var port int
var server structures.Server

// читаем из .енв все значения
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
	server := structures.StartServer(structures.HandleMessage, port, config.WIDTH, config.HEIGHT, config.G)
	log.Println("Started server on", server.Port)
	for {
		continue
	}
}
