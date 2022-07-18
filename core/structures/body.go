package core

import "math"

type Body struct {
	id      uint16
	X, Y    int
	mass    float64
	speedUp Vector
	speed   Vector
	factor  *float64
}

func (b Body) Coords() (int, int) {
	return b.X, b.Y
}

func (b Body) Mass() float64 {
	return b.mass
}

func (b Body) updateCoords() {
	var speedX float64 = b.speed.X
	var speedY float64 = b.speed.Y

	b.X += int(math.Round(speedX * *b.factor))
	b.Y += int(math.Round(speedY * *b.factor))
}

func (b Body) updateSpeedup(force *Vector) {
	var dSpeedup Vector = force.div(b.mass)
	b.speedUp = b.speedUp.add(&dSpeedup)
}

func (b Body) updateSpeed() {
	var dSpeed Vector = b.speedUp.mul(*b.factor)
	b.speed = b.speed.add(&dSpeed)
}

func (b Body) updateAll(force *Vector) {
	b.updateSpeedup(force)
	b.updateSpeed()
	b.updateCoords()
}
