package structures

import (
	"encoding/json"
	core "github.com/migelit0/physics_server/core/structures"
	"log"
)

type Response struct {
	Data map[uint16][2]int
}

func generateResponse(w *core.World) Response {
	var res Response
	var coords [2]int
	for _, body := range w.Bodies {
		coords = [2]int{body.X, body.Y}
		res.Data[body.Id] = coords
	}

	return res
}

func generateResponseText(w *core.World) []byte {
	var res = generateResponse(w)
	log.Println(res.Data)

	jsonData, err := json.Marshal(res.Data)
	if err != nil {
		log.Panicln(err)
		return []byte{}
	}

	return jsonData
}
