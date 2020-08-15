package quicksort

import "github.com/tolbier/algorithms_go/courses/course1/week3/quicksort/pivoter"

type integers []int

func (ints integers) QuickSort() int {
	return ints.PivotSort(pivoter.NewRandomPivoter())
}
func (ints integers) PivotSort(pivoter pivoter.Pivoter) int {
	if len(ints) <= 1 {
		return 0
	}
	i := pivoter.Pivot(ints)
	ints[0], ints[i] = ints[i], ints[0]
	l, r, comparisons := ints.partition()
	comparisons += l.PivotSort(pivoter)
	comparisons += r.PivotSort(pivoter)

	return comparisons
}
func (ints integers) partition() (l, r integers, comparisons int) {
	comparisons = len(ints) - 1
	p := ints[0]
	i := 1
	for j := 1; j < len(ints); j++ {
		if ints[j] < p {
			ints[i], ints[j] = ints[j], ints[i]
			i++
		}
	}
	ints[0], ints[i-1] = ints[i-1], ints[0]
	l, r = ints[:i-1], ints[i:]
	return
}
