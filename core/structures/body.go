package core

import "math"

type Body struct {
	Id         uint16
	X, Y       int
	Mass       float64
	SpeedUp    Vector
	Speed      Vector
	Factor     *float64
	MaxX, MaxY *uint16
}

func (b Body) Coords() (int, int) {
	return b.X, b.Y
}

func (b *Body) updateCoords() {
	var speedX float64 = b.Speed.X
	var speedY float64 = b.Speed.Y

	b.X += int(math.Round(speedX * *b.Factor))
	b.Y += int(math.Round(speedY * *b.Factor))
}

func (b *Body) updateSpeedup(force *Vector) {
	var newSpeedup Vector = force.div(b.Mass)
	b.SpeedUp = newSpeedup
}

func (b *Body) updateSpeed() {
	var dSpeed Vector = b.SpeedUp.mul(*b.Factor)
	b.Speed = b.Speed.add(&dSpeed)
}

func (b *Body) updateAll(force *Vector) {
	b.updateSpeedup(force)
	b.updateSpeed()
	b.updateCoords()
}

func (b *Body) correctCoords() {
	if b.X <= 0 || b.X >= *b.MaxX {
		b.Speed.setX(-1 * b.Speed.X)
		if b.X <= 0 {
			b.X = 1
		}
		if b.X >= *b.MaxX {
			b.Y = *b.MaxX - 1
		}
	}

	if b.Y <= 0 || b.Y >= *b.MaxY {
		b.Speed.setY(-1 * b.Speed.Y)
		if b.Y <= 0 {
			b.Y = 1
		}
		if b.Y >= *b.MaxY {
			b.Y = *b.MaxY - 1
		}
	}
}
