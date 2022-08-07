package structures

import (
	"log"
)

func HandleMessage(message []byte) {
	log.Println("get", string(message))
}
