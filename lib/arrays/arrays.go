package arrays

func SplitHalves(numbers []int) (h1, h2 []int) {
	n := len(numbers)
	hn := n / 2
	return numbers[:hn], numbers[hn:]
}
