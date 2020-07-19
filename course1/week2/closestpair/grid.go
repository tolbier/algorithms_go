package closestpair

import (
	"sort"

	"github.com/tolbier/algorithms_go/course1/week2/closestpair/gridsort"
	"github.com/tolbier/algorithms_go/course1/week2/closestpair/model"
)

type Grid []*model.Point

func (g Grid) ClosestPair() *model.PointPair {
	return NewSortedGrid(g).ClosestPair()
}

func NewSortedGrid(g Grid) *SortedGrid {
	return NewSortedGridFromGridSorted(g.SortBy(gridsort.X), g.SortBy(gridsort.Y))
}

func NewSortedGridFromGridSorted(gSorted Grid, oldGy Grid) *SortedGrid {
	sortedGrid := SortedGrid{gx: gSorted}
	sortedGrid.createGridMap()
	sortedGrid.createYGrid(oldGy)
	return &sortedGrid
}

func (sg *SortedGrid) createYGrid(oldGy Grid) {
	sg.gy = Grid{}
	for _, p := range oldGy {
		if sg.gMap.hasPoint(*p) {
			sg.gy = append(sg.gy, p)
		}
	}
}

type GridMap map[model.Point]bool
type SortedGrid struct {
	gx, gy Grid
	gMap   GridMap
}

func (sg *SortedGrid) createGridMap() {
	sg.gMap = GridMap{}
	for _, v := range sg.gx {
		sg.gMap[*v] = true
	}
}

func (gm *GridMap) hasPoint(p model.Point) bool {
	_, ok := (*gm)[p]
	return ok
}

func (g Grid) SortBy(f func(i, j *model.Point) bool) (res Grid) {
	res = make(Grid, len(g))
	copy(res, g)
	sort.Sort(gridsort.New(res, f))
	return res
}

func (sg *SortedGrid) Len() int {
	return len(sg.gx)
}

func (sg *SortedGrid) SplitHalves() (sg1 *SortedGrid, sg2 *SortedGrid) {
	hx1, hx2 := SplitHalves(sg.gx)
	sg1 = NewSortedGridFromGridSorted(hx1, sg.gy)
	sg2 = NewSortedGridFromGridSorted(hx2, sg.gy)

	return
}
func (g Grid) last() *model.Point {
	return g[len(g)-1]
}

func (sg *SortedGrid) lastX() float64 {
	return sg.gx.last().X
}
func (sg *SortedGrid) ClosestPair() *model.PointPair {
	if sg.Len() <= 3 {
		return sg.ClosestPairBasicCase()
	}
	lx, rx := sg.SplitHalves()
	lp := lx.ClosestPair()
	rp := rx.ClosestPair()
	bp := BestPair(lp, rp)
	s := sg.ClosestSplitPair(lx.lastX(), bp.Dist())

	return BestPair(bp, s)
}

func (sg *SortedGrid) ClosestPairBasicCase() (res *model.PointPair) {
	gx := sg.gx
	a := &model.PointPair{A: gx[0], B: gx[1]}
	if sg.Len() == 2 {
		return a
	}
	b := &model.PointPair{A: gx[0], B: gx[2]}
	c := &model.PointPair{A: gx[1], B: gx[2]}
	return BestPair(BestPair(a, b), c)
}

func BestPair(p, q *model.PointPair) (bp *model.PointPair) {
	if p != nil && (q == nil || p.Dist() <= q.Dist()) {
		return p
	}
	return q
}

func (sg *SortedGrid) ClosestSplitPair(xMedian float64, minDist float64) (pp *model.PointPair) {
	var sy Grid
	for _, q := range sg.gy {
		if q.X > xMedian-minDist && q.X < xMedian+minDist {
			sy = append(sy, q)
		}
	}
	for i := 0; i < len(sy)-1; i++ {
		v := sy[i]
		for j := 1; j+i < len(sy) && j <= 7; j++ {
			w := sy[i+j]
			if v.Dist(w) < minDist {
				minDist = v.Dist(w)
				pp = &model.PointPair{A: v, B: w}
			}
		}
	}
	return
}
