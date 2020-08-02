package quicksort

type integers []int

func (ints integers) QuickSort(pivoter Pivoter) int {
	if len(ints) <= 1 {
		return 0
	}
	i := pivoter.Pivot(ints)
	ints[0], ints[i] = ints[i], ints[0]
	l, r := ints.partition()
	comparisons := len(ints) - 1
	comparisons += l.QuickSort(pivoter)
	comparisons += r.QuickSort(pivoter)

	return comparisons
}
func (ints integers) partition() (l, r integers) {
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
