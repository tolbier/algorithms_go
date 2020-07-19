package closestpair

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/course1/week2/closestpair/model"
)

func Test_Closest_Pair(t *testing.T) {
	grid := Grid{
		{1, 6},
		{3, 2},
		{1, 8.1},
		{1, 10},
		{1, 4},
		{1, 12},
		{1, 8},
		{1, 2},
	}
	p := grid.ClosestPair() //return a PointPair
	assert.True(t,
		p.Is(model.Point{1, 8}, model.Point{1, 8.1}))
}
func Test_Closest_Pair2(t *testing.T) {
	grid := Grid{
		{1, 2},
		{1, 4},
		{1, 6},
		{1, 8},
		{1, 9.9},
		{1, 10},
		{1, 12},
	}
	p := grid.ClosestPair()
	assert.True(t,
		p.Is(model.Point{1, 9.9}, model.Point{1, 10}))
}
