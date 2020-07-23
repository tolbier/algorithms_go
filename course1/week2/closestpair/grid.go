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

func (g Grid) SortBy(f func(i, j *model.Point) bool) (res Grid) {
	res = make(Grid, len(g))
	copy(res, g)
	sort.Sort(gridsort.New(res, f))
	return res
}

func (g Grid) last() *model.Point {
	return g[len(g)-1]
}
func (g Grid) SplitHalves() (h1, h2 []*model.Point) {
	n := len(g)
	hn := n / 2
	return g[:hn], g[hn:]
}
