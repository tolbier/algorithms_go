package quicksort

type Pivoter interface {
	Pivot(list []int) (pivot int)
}
