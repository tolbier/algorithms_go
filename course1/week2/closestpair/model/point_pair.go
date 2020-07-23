package model

type PointPair struct {
	a, b *Point
}

func NewPointPair(a, b *Point) *PointPair {
	return &PointPair{
		a: a,
		b: b,
	}
}
func (pp *PointPair) Dist() float64 {
	return pp.a.Dist(pp.b)
}

func (pp *PointPair) Is(p1, p2 Point) bool {
	return (*pp.a == p1 && *pp.b == p2) || (*pp.a == p2 && *pp.b == p1)
}
func (pp *PointPair) BestPair(pq *PointPair) (bp *PointPair) {
	if pq == nil || pp.Dist() <= pq.Dist() {
		return pp
	}
	return pq
}
