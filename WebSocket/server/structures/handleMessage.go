package structures

import core "github.com/migelit0/physics_server/core/structures"

func handleMessage(message []byte) {

}

func generateResponse(w *core.World) Response {
	var res Response
	var coords [2]int
	for i, body := range w.Bodies {
		coords = [2]int{body.X, body.Y}
		res.Data[body.Id] = coords
	}

}
