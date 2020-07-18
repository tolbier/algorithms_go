package closestpair

import (
	"math"
	"sort"
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
type GridSort struct {
	g    Grid
	less func(i, j Point) bool
}

func (gs GridSort) Len() int           { return len(gs.g) }
func (gs GridSort) Swap(i, j int)      { gs.g[i], gs.g[j] = gs.g[j], gs.g[i] }
func (gs GridSort) Less(i, j int) bool { return gs.less(gs.g[i], gs.g[j]) }

func X(i, j Point) bool { return i.X < j.X }
func Y(i, j Point) bool { return i.Y < j.Y }

type PointPair struct {
	a, b Point
}

func (g Grid) ClosestPair() PointPair {
	return g.SortBy(X).ClosestPair()

}
func (g Grid) SortBy(f func(i, j Point) bool) (res SortedGrid) {
	res = make(SortedGrid, len(g))
	copy(res, g)
	sort.Sort(GridSort{Grid(res), f})
	return res
}

func (g SortedGrid) Len() int {
	return len(g)
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
	s := ClosestSplitPair(sg, Grid(sg).SortBy(Y), minDist)
	if s == nil {
		return bp
	}
	return BestPair(bp, *s)
}

func ClosestPairBasicCase(px SortedGrid) (res PointPair) {
	minDist := math.MaxFloat64
	for i, v := range px {
		for j := 1; j+i < len(px); j++ {
			w := px[i+j]
			pp := PointPair{v, w}
			if pp.dist() < minDist {
				minDist = pp.dist()
				res = pp
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