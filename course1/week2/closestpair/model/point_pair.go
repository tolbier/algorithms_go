package model

type PointPair struct {
	A, B *Point
}

func (pp PointPair) Dist() float64 {
	return pp.A.Dist(pp.B)
}

func (pp PointPair) Is(p1, p2 Point) bool {
	return (*pp.A == p1 && *pp.B == p2) || (*pp.A == p2 && *pp.B == p2)
}
