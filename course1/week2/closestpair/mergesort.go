package closestpair

import "github.com/tolbier/algorithms_go/course1/week2/closestpair/model"

func SplitHalves(numbers []*model.Point) (h1, h2 []*model.Point) {
	n := len(numbers)
	hn := n / 2
	return numbers[:hn], numbers[hn:]
}
