package closestpair

import "github.com/tolbier/algorithms_go/course1/week2/closestpair/model"

type GridMap map[*model.Point]bool

func (gm *GridMap) hasPoint(p *model.Point) bool {
	_, ok := (*gm)[p]
	return ok
}
