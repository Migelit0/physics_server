package main

import core "github.com/migelit0/physics_server/core/structures"

func initWorld() core.World {
	var emptyBodies []core.Body
	w := core.World{Width: WIDTH, Height: HEIGHT, Bodies: emptyBodies, G: &G}
	return w
}

func main() {

}
