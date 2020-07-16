package closestpair

import (
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) dist(q Point) float64 {
	return math.Sqrt(math.Pow(p.X-q.X, 2) + math.Pow(p.Y-q.Y, 2))
}

func (pp PointPair) dist() float64 {
	return pp.a.dist(pp.b)
}

func (pp PointPair) is(pp2 PointPair) bool {
	return (pp.a == pp2.a && pp.b == pp2.b) || (pp.a == pp2.b && pp.b == pp2.a)
}

type Grid []Point
type SortedGrid Grid
type PointPair struct {
	a, b Point
}

func (g Grid) ClosestPair() PointPair {
	return g.SortByX().ClosestPair()

}
func (g Grid) SortByX() SortedGrid {
	return Mergesort(g, func(a, b Point) bool {
		return a.X <= b.X
	})
}

func (g SortedGrid) SortByY() SortedGrid {
	return Mergesort(g, func(a, b Point) bool {
		return a.Y <= b.Y
	})
}
func (g SortedGrid) SplitHalves() (SortedGrid, SortedGrid) {
	return SplitHalves(g)

}
func (g SortedGrid) lastPoint() Point {
	return g[len(g)-1]
}
func (sg SortedGrid) ClosestPair() PointPair {
	if len(sg) <= 3 {
		return ClosestPairBasicCase(sg)
	}
	lx, rx := sg.SplitHalves()
	l := lx.ClosestPair()
	r := rx.ClosestPair()
	bp := BestPair(l, r)
	minDist := bp.dist()
	s := ClosestSplitPair(sg, sg.SortByY(), minDist)
	if s == nil {
		return bp
	}
	return BestPair(bp, *s)
}

func ClosestPairBasicCase(px SortedGrid) (pp PointPair) {
	minDist := math.MaxFloat64
	for i, v := range px {
		for j := 1; j+i < len(px); j++ {
			w := px[i+j]
			if v.dist(w) < minDist {
				minDist = v.dist(w)
				pp = PointPair{v, w}
			}
		}
	}
	return
}
func BestPair(p, q PointPair) (bp PointPair) {
	if p.dist() <= q.dist() {
		return p
	}
	return q
}
func ClosestSplitPair(px, py SortedGrid, minDist float64) (pp *PointPair) {
	lx, _ := px.SplitHalves()
	xMedian := lx.lastPoint().X
	var sy Grid
	for _, q := range py {
		if q.X > xMedian-minDist && q.X < xMedian+minDist {
			sy = append(sy, q)
		}
	}
	for i := 0; i < len(sy)-1; i++ {
		v := sy[i]
		for j := 1; j+i < len(sy) && j <= 7; j++ {
			w := sy[i+j]
			if v.dist(w) < minDist {
				minDist = v.dist(w)
				pp = &PointPair{v, w}
			}
		}
	}
	return
}
