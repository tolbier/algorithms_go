package quicksort

type firstPivoter struct{}

func NewFirstPivoter() Pivoter {
	return &firstPivoter{}
}
func (fp *firstPivoter) Pivot(_ []int) (pivot int) {
	return 0
}
