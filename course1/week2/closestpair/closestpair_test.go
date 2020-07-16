package closestpair

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		p.is(PointPair{
			Point{1, 8}, Point{1, 8.1}}))
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
		p.is(PointPair{
			Point{1, 9.9}, Point{1, 10}}))
}
