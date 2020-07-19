package model

import "math"

type Point struct {
	X, Y float64
}

func (p *Point) Dist(q *Point) float64 {
	return math.Sqrt(math.Pow(p.X-q.X, 2) + math.Pow(p.Y-q.Y, 2))
}
