package secondlargest

import (
	"github.com/tolbier/algorithms_go/course1/week1/mergesort"
	"github.com/tolbier/algorithms_go/lib/arrays"
)

func SecondLargest(numbers []int) (res int) {
	res = TwoLargest(numbers)[0]
	return
}

func TwoLargest(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	h1, h2 := arrays.SplitHalves(numbers)
	l1 := TwoLargest(h1)
	l2 := TwoLargest(h2)
	mrg := mergesort.Merge(l1, l2)
	return mrg[len(mrg)-2:]
}
