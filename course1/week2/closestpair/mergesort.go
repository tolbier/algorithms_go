package closestpair

func SplitHalves(numbers []*Point) (h1, h2 []*Point) {
	n := len(numbers)
	hn := n / 2
	return numbers[:hn], numbers[hn:]
}
