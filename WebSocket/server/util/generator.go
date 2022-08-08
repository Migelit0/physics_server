package util

import (
	"github.com/migelit0/physics_server/WebSocket/server/config"
	core "github.com/migelit0/physics_server/core/structures"
	"math/rand"
	"time"
)

func GenerateInitWorld(amount int) core.World {
	var bodies []core.Body

	g := config.G
	width := config.WIDTH
	height := config.HEIGHT
	factor := config.FACTOR

	var newWorld core.World = core.World{
		Width:  config.WIDTH,
		Height: config.HEIGHT,
		Bodies: bodies,
		G:      &g,
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < amount; i++ {

		x := rand.Intn(int(config.WIDTH))
		y := rand.Intn(int(config.HEIGHT))

		body := core.Body{
			Id:      uint16(i),
			X:       x,
			Y:       y,
			Mass:    1e14,
			SpeedUp: core.Vector{},
			Speed:   core.Vector{},
			Factor:  &factor,
			MaxX:    &width,
			MaxY:    &height,
		}
		newWorld.AppendBody(&body)
	}

	return newWorld
}
