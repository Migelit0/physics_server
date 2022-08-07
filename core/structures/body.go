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
	var speedX = b.Speed.X
	var speedY = b.Speed.Y

	b.X += int(math.Round(speedX * *b.Factor))
	b.Y += int(math.Round(speedY * *b.Factor))
}

func (b *Body) updateSpeedup(force *Vector) {
	var newSpeedup = force.Div(&b.Mass)
	b.SpeedUp = newSpeedup
}

func (b *Body) updateSpeed() {
	var dSpeed = b.SpeedUp.Mul(b.Factor)
	b.Speed = b.Speed.Add(&dSpeed)
}

func (b *Body) updateAll(force *Vector) {
	b.updateSpeedup(force)
	b.updateSpeed()
	b.updateCoords()
}

func (b *Body) correctCoords() {
	if b.X <= 0 || b.X >= int(*b.MaxX) {
		b.Speed.SetX(-1 * b.Speed.X)
		if b.X <= 0 {
			b.X = 1
		}
		if b.X >= int(*b.MaxX) {
			b.Y = int(*b.MaxX) - 1
		}
	}

	if b.Y <= 0 || b.Y >= int(*b.MaxY) {
		b.Speed.SetY(-1 * b.Speed.Y)
		if b.Y <= 0 {
			b.Y = 1
		}
		if b.Y >= int(*b.MaxY) {
			b.Y = int(*b.MaxY) - 1
		}
	}
}
