package secondlargest

import "github.com/tolbier/algorithms_go/course1/week1/mergesort"

func SecondLargest(numbers []int) (res int) {
	res = TwoLargest(numbers)[0]
	return
}

func TwoLargest(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}
	h1, h2 := splitHalves(numbers)
	l1 := TwoLargest(h1)
	l2 := TwoLargest(h2)
	mrg := mergesort.Merge(l1, l2)
	return mrg[len(mrg)-2:]
}
func splitHalves(numbers []int) (h1, h2 []int) {
	n := len(numbers)
	hn := n / 2
	return numbers[:hn], numbers[hn:]
}
