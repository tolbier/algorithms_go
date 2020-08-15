package inversions

import "github.com/tolbier/algorithms_go/lib/arrays"

func Count(numbers []int) (count int) {
	count, _ = CountMerge(numbers)
	return
}
func CountMerge(numbers []int) (count int, merge []int) {
	if len(numbers) == 1 {
		return 0, numbers
	}
	h1, h2 := arrays.SplitHalves(numbers)
	i1, hs1 := CountMerge(h1)
	i2, hs2 := CountMerge(h2)
	count, merge = CountMergeSplit(hs1, hs2)
	count += i1 + i2
	return
}

func CountMergeSplit(hs1 []int, hs2 []int) (count int, merge []int) {
	idx1, idx2 := 0, 0
	n1 := len(hs1)
	n2 := len(hs2)
	n := n1 + n2
	for idx1+idx2 < n {
		if idx2 >= n2 || (idx1 < n1 && hs1[idx1] <= hs2[idx2]) {
			merge = append(merge, hs1[idx1])
			idx1++
		} else {
			merge = append(merge, hs2[idx2])
			count += n1 - idx1

			idx2++
		}
	}
	return
}
