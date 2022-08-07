package core

type Vector struct {
	X, Y float64
}

// Add two vectors
func (v Vector) Add(v2 *Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

// Mul Multiply vector and coefficient
func (v Vector) Mul(k float64) Vector {
	return Vector{v.X * k, v.Y * k}
}

// Div Divide vector and number
func (v Vector) Div(k float64) Vector {
	return Vector{v.X / k, v.Y / k}
}

func (v *Vector) SetX(newX float64) {
	v.X = newX
}

func (v *Vector) SetY(newY float64) {
	v.Y = newY
}
