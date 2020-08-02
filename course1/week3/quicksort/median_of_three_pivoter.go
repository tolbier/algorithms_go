package quicksort

import (
	"sort"
)

type medianOfThreePivoter struct{}

func NewMedianOfThreePivoter() Pivoter {
	return &medianOfThreePivoter{}
}

func (fp *medianOfThreePivoter) Pivot(list []int) (pivot int) {
	l, r := 0, len(list)-1
	m := (r + l) / 2
	vMap := map[int]int{
		list[l]: l,
		list[r]: r,
		list[m]: m,
	}
	keys := []int{list[l], list[r], list[m]}
	sort.Ints(keys)
	pivot = vMap[keys[1]]
	return
}
