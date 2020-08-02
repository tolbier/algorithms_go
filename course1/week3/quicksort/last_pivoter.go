package quicksort

type lastPivoter struct{}

func NewLastPivoter() Pivoter {
	return &lastPivoter{}
}

func (fp *lastPivoter) Pivot(list []int) (pivot int) {
	return len(list) - 1
}
