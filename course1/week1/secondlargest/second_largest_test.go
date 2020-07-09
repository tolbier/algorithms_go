package secondlargest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Second_Largest(t *testing.T) {
	example_arrays := [...][]int{
		{1, 2, 3, 4, 5},
		{5, 7, 9, 3, 8},
		{1, 3, 5, 2, 7},
		{2, 4, 6, 8},
		{1, 0, 4, -6},
		{11, 12, -13, 14, 15},
		{2, 1},
		{-210, 0},
	}
	example_sols := [...]int{4, 8, 5, 6, 1, 14, 1, -210}
	for i, seq := range example_arrays {
		sol := example_sols[i]
		secLargest := SecondLargest(seq)
		assert.Equal(t, sol, secLargest)

	}
}

func Test_Two_Largest2(t *testing.T) {
	a := SecondLargest([]int{3, 4, 5})
	assert.Equal(t, 4, a)
}
