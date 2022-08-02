package core

import "math"

type World struct {
	Width, Height uint16
	Bodies        []Body
	G             *float64
}

// AppendBody spawn new body in world
func (w *World) AppendBody(body *Body) {
	w.Bodies = append(w.Bodies, *body)
}

// CalcForceForBody Calculate force for body with index
func (w *World) CalcForceForBody(index *int) Vector {
	var resultantForce = Vector{0, 0}
	var force Vector
	var bodyMain = w.Bodies[*index]

	for i, body := range w.Bodies {
		if i == *index {
			continue
		}
		force = w.calcForceTwoBodies(&bodyMain, &body)
		resultantForce = resultantForce.add(&force)
	}
	return resultantForce
}

func (w *World) CalcForceTwoBodies(b0, b1 *Body) Vector {
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

func (w *World) calcAbcForceTwoBodies(b0, b1 *Body, R float64) float64 {
	return *w.G * b0.Mass * b1.Mass / (R * R)
}

func (w *World) handleBody(index *int) {
	var force Vector
	force = w.CalcForceForBody(index)
	w.Bodies[*index].updateSpeedup(&force)
	w.Bodies[*index].updateSpeed()
}

func (w *World) DoOneIter() {
	for i := range w.Bodies {
		// считаем силу, ускорение, скорость для каждого тела
		w.handleBody(&i)
	}

	for i := range w.Bodies {
		// обновляем для всех тел координаты
		w.Bodies[i].updateCoords()
	}

	for i := range w.Bodies {
		// проверяем валидны ли координаты и обрабатываем отскоки
		w.Bodies[i].correctCoords()
	}
}
