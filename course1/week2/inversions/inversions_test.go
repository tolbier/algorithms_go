package inversions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/course1/week2/inversions"
)

func Test_Inversions(t *testing.T) {
	a := inversions.Count([]int{1, 3, 5, 2, 4, 6})
	assert.Equal(t, 3, a)
}

func Test_Inversions2(t *testing.T) {
	a := inversions.Count([]int{6, 5, 4, 3, 2, 1})
	assert.Equal(t, 15, a)
}

//func Test_Merge2(t *testing.T) {
//	a := mergesort.Merge(
//		[]int{4, 5},
//		[]int{6, 7, 8})
//	assert.Equal(t, []int{4, 5, 6, 7, 8}, a)
//}
//func Test_MergeSort(t *testing.T) {
//	a := mergesort.Mergesort([]int{3, 54, 56, 67, 7, 8, 9, 1, 3, 4, 5})
//	assert.Equal(t, []int{1, 3, 3, 4, 5, 7, 8, 9, 54, 56, 67}, a)
//}
