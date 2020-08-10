package rselect

import "github.com/tolbier/algorithms_go/course1/week3/quicksort/pivoter"

type integers []int

func (ints integers) RSelect(i int) int {
	return ints.PivotRSelect(i, pivoter.NewRandomPivoter())
}

func (ints integers) PivotRSelect(i int, pivoter pivoter.Pivoter) int {
	if len(ints) == 1 {
		return ints[0]
	}
	p := pivoter.Pivot(ints)
	ints[0], ints[p] = ints[p], ints[0]
	j := ints.partitionPosition()
	if j == i {
		return ints[i]
	} else if j > i {
		return ints[:j].PivotRSelect(i, pivoter)
	}
	return ints[j+1:].PivotRSelect(i-(j+1), pivoter)
}
func (ints integers) partitionPosition() int {
	p := ints[0]
	i := 1
	for j := 1; j < len(ints); j++ {
		if ints[j] < p {
			ints[i], ints[j] = ints[j], ints[i]
			i++
		}
	}
	ints[0], ints[i-1] = ints[i-1], ints[0]
	return i - 1
}
