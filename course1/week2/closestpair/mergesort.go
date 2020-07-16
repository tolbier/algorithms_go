package closestpair

func Mergesort(numbers []Point, isBefore func(a, b Point) bool) []Point {
	if len(numbers) == 1 {
		return numbers
	}
	h1, h2 := SplitHalves(numbers)
	hs1 := Mergesort(h1, isBefore)
	hs2 := Mergesort(h2, isBefore)
	return Merge(hs1, hs2, isBefore)
}

func Merge(hs1 []Point, hs2 []Point, isBefore func(a, b Point) bool) (res []Point) {
	idx1, idx2 := 0, 0
	n1 := len(hs1)
	n2 := len(hs2)
	n := n1 + n2
	for idx1+idx2 < n {
		if idx2 >= n2 || (idx1 < n1 && isBefore(hs1[idx1], hs2[idx2])) {
			res = append(res, hs1[idx1])
			idx1++
		} else {
			res = append(res, hs2[idx2])
			idx2++
		}
	}
	return
}
func SplitHalves(numbers []Point) (h1, h2 []Point) {
	n := len(numbers)
	hn := n / 2
	return numbers[:hn], numbers[hn:]
}
