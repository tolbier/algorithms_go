package mergesort

import "github.com/tolbier/algorithms_go/lib/arrays"

func Mergesort(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	}
	h1, h2 := arrays.SplitHalves(numbers)
	hs1 := Mergesort(h1)
	hs2 := Mergesort(h2)
	return Merge(hs1, hs2)
}

func Merge(hs1 []int, hs2 []int) (res []int) {
	idx1, idx2 := 0, 0
	n1 := len(hs1)
	n2 := len(hs2)
	n := n1 + n2
	for idx1+idx2 < n {
		if idx2 >= n2 || (idx1 < n1 && hs1[idx1] <= hs2[idx2]) {
			res = append(res, hs1[idx1])
			idx1++
		} else {
			res = append(res, hs2[idx2])
			idx2++
		}
	}
	return
}
