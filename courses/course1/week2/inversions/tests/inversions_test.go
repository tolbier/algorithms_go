package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/courses/course1/week2/inversions"
	"github.com/tolbier/algorithms_go/lib/testcases"
)

func Test_InversionsCount(t *testing.T) {
	a := inversions.Count([]int{1, 3, 5, 2, 4, 6})
	assert.Equal(t, 3, a)
}

func Test_InversionsCount2(t *testing.T) {
	a := inversions.Count([]int{6, 5, 4, 3, 2, 1})
	assert.Equal(t, 15, a)
}
func Test_Inversions_Count_TestCases(t *testing.T) {
	inputs, outputs := testcases.InputsOutputs("testcases")
	for idx, input := range inputs {
		inputInts := testcases.ParseIntegerFile(input)
		outputInts := testcases.ParseIntegerFile(outputs[idx])
		result := inversions.Count(inputInts)
		assert.Equal(t, result, outputInts[0])
	}
}
