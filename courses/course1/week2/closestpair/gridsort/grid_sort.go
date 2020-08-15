package gridsort

import (
	"github.com/tolbier/algorithms_go/courses/course1/week2/closestpair/model"
)

type GridSort struct {
	g    []*model.Point
	less func(i, j *model.Point) bool
}

func New(g []*model.Point, f func(i, j *model.Point) bool) *GridSort {
	return &GridSort{
		g:    g,
		less: f,
	}
}
func (gs *GridSort) Len() int           { return len(gs.g) }
func (gs *GridSort) Swap(i, j int)      { gs.g[i], gs.g[j] = gs.g[j], gs.g[i] }
func (gs *GridSort) Less(i, j int) bool { return gs.less(gs.g[i], gs.g[j]) }

func X(i, j *model.Point) bool { return i.X < j.X }
func Y(i, j *model.Point) bool { return i.Y < j.Y }
