package closestpair

import (
	"github.com/tolbier/algorithms_go/courses/course1/week2/closestpair/gridsort"
	"github.com/tolbier/algorithms_go/courses/course1/week2/closestpair/model"
)

type SortedGrid struct {
	gx, gy Grid
	gMap   GridMap
	ch     chan *model.PointPair
}

func NewSortedGrid(g Grid) *SortedGrid {
	return NewSortedGridFromGrids(g.SortBy(gridsort.X), g.SortBy(gridsort.Y))
}

func NewSortedGridFromGrids(gSorted Grid, oldGy Grid) *SortedGrid {
	sortedGrid := SortedGrid{gx: gSorted, ch: make(chan *model.PointPair)}
	sortedGrid.createGridMap()
	sortedGrid.createYGrid(oldGy)
	return &sortedGrid
}

func (sg *SortedGrid) createGridMap() {
	sg.gMap = GridMap{}
	for _, v := range sg.gx {
		sg.gMap[v] = true
	}
}
func (sg *SortedGrid) createYGrid(oldGy Grid) {
	sg.gy = Grid{}
	for _, p := range oldGy {
		if sg.gMap.hasPoint(p) {
			sg.gy = append(sg.gy, p)
		}
	}
}
func (sg *SortedGrid) Len() int {
	return len(sg.gx)
}
func (sg *SortedGrid) SplitHalves() (sg1 *SortedGrid, sg2 *SortedGrid) {
	hx1, hx2 := sg.gx.SplitHalves()
	sg1 = NewSortedGridFromGrids(hx1, sg.gy)
	sg2 = NewSortedGridFromGrids(hx2, sg.gy)

	return
}

func (sg *SortedGrid) lastX() float64 {
	return sg.gx.last().X
}

func (sg *SortedGrid) ClosestPairRun() {
	sg.ch <- sg.ClosestPair()
}
func (sg *SortedGrid) ClosestPair() *model.PointPair {
	if sg.Len() <= 3 {
		return sg.ClosestPairBasicCase()
	}
	lx, rx := sg.SplitHalves()
	go lx.ClosestPairRun()
	go rx.ClosestPairRun()

	lp, rp := <-lx.ch, <-rx.ch

	bp := lp.BestPair(rp)
	s := sg.ClosestSplitPair(lx.lastX(), bp.Dist())

	return bp.BestPair(s)
}

func (sg *SortedGrid) ClosestPairBasicCase() (res *model.PointPair) {
	gx := sg.gx
	a := gx[0].Pair(gx[1])
	if sg.Len() == 2 {
		return a
	}
	b := gx[0].Pair(gx[2])
	c := gx[1].Pair(gx[2])
	return a.BestPair(b).BestPair(c)
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
				pp = v.Pair(w)
			}
		}
	}
	return
}
