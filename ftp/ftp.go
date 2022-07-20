package main

import (
	"fmt"
	"github.com/joho/godotenv"
	core "github.com/migelit0/physics_server/core/structures"
	"github.com/migelit0/physics_server/ftp/config"
	"github.com/migelit0/physics_server/ftp/ftp"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

var port int
var rootDir string

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

	rootDirEnv, isExist := os.LookupEnv("ROOT_DIR")
	if isExist {
		rootDir = rootDirEnv
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
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatal(err)
	}
	ftp.Serve(ftp.NewConn(c, absPath))
}
