package quicksort

func QuickSort(slc []int, pivoter Pivoter) int {
	if len(slc) <= 1 {
		return 0
	}
	i := pivoter.Pivot(slc)
	slc[0], slc[i] = slc[i], slc[0]
	l, r := partition(slc)
	comparisons := len(slc) - 1
	comparisons += QuickSort(l, pivoter)
	comparisons += QuickSort(r, pivoter)

	return comparisons
}
func partition(slc []int) (l, r []int) {
	p := slc[0]
	i := 1
	for j := 1; j < len(slc); j++ {
		if slc[j] < p {
			slc[i], slc[j] = slc[j], slc[i]
			i++
		}
	}
	slc[0], slc[i-1] = slc[i-1], slc[0]
	l, r = slc[:i-1], slc[i:]
	return
}
