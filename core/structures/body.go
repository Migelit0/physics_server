package core

import "math"

type Body struct {
	id         uint16
	X, Y       int
	mass       float64
	speedUp    Vector
	speed      Vector
	factor     *float64
	maxX, maxY *int
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
	var newSpeedup Vector = force.div(b.mass)
	b.speedUp = newSpeedup
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

func (b Body) correctCoords() {
	if b.X <= 0 || b.X >= *b.maxX {
		b.speed.setX(-1 * b.speed.X)
		if b.X <= 0 {
			b.X = 1
		}
		if b.X >= *b.maxX {
			b.Y = *b.maxX - 1
		}
	}

	if b.Y <= 0 || b.Y >= *b.maxY {
		b.speed.setY(-1 * b.speed.Y)
		if b.Y <= 0 {
			b.Y = 1
		}
		if b.Y >= *b.maxY {
			b.Y = *b.maxY - 1
		}
	}
}
