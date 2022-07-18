package core

import "math"

type World struct {
	width, height int16
	bodies        []Body
	G             *float64
}

func (w World) appendBody(body *Body) {
	w.bodies = append(w.bodies, *body)
}

// Calculate force for body with index
func (w World) calcForceForBody(index *int) Vector {
	var resultantForce Vector = Vector{0, 0}
	var force Vector
	var bodyMain Body = w.bodies[*index]

	for i, body := range w.bodies {
		if i == *index {
			continue
		}
		force = w.calcForceTwoBodies(&bodyMain, &body)
		resultantForce = resultantForce.add(&force)
	}
	return resultantForce
}

func (w World) calcForceTwoBodies(b0, b1 *Body) Vector {
	var x1, y1, x, y, dx, dy int
	var dist, abcForce, sin, cos float64
	var force Vector

	force = Vector{0, 0}
	x1, y1 = b0.Coords()
	x, y = b1.Coords()

	dx = x1 - x
	dy = y1 - y
	dist = math.Sqrt(float64(dx*dx + dy*dy))
	sin = float64(dy) / dist
	cos = float64(dx) / dist

	abcForce = w.calcAbcForceTwoBodies(b0, b1, dist)
	force = Vector{abcForce * cos, abcForce * sin}

	return force
}

func (w World) calcAbcForceTwoBodies(b0, b1 *Body, R float64) float64 {
	return *w.G * b0.Mass() * b1.Mass() / (R * R)
}

func (w World) handleBody(index *int) {
	var force Vector
	force = w.calcForceForBody(index)
	w.bodies[*index].updateSpeedup(&force)
	w.bodies[*index].updateSpeed()
}

func (w World) doOneIter() {
	for i, _ := range w.bodies {
		w.handleBody(&i)
	}

	for i, _ := range w.bodies {
		w.bodies[i].updateCoords()
	}
}
