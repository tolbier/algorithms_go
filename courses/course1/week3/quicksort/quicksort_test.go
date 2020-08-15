package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/courses/course1/week3/quicksort/pivoter"
	"github.com/tolbier/algorithms_go/lib/testcases"
)

func Test_PivotSort_TestCases(t *testing.T) {
	inputs, outputs := testcases.InputsOutputs("testcases")
	for idx, input := range inputs {
		inputInts := testcases.ParseIntegerFile(input)
		outputInts := testcases.ParseIntegerFile(outputs[idx])
		result := ThreePivotSorts(inputInts)
		assert.Equal(t, result, outputInts)
	}
}

func ThreePivotSorts(inputInts []int) (res []int) {
	slc := make(integers, len(inputInts))
	copy(slc, inputInts)
	res = append(res, slc.PivotSort(pivoter.NewFirstPivoter()))
	copy(slc, inputInts)
	res = append(res, slc.PivotSort(pivoter.NewLastPivoter()))
	copy(slc, inputInts)
	res = append(res, slc.PivotSort(pivoter.NewMedianOfThreePivoter()))
	return
}

func Test_QuickSort_TestCases(t *testing.T) {
	inputs, _ := testcases.InputsOutputs("testcases")
	for _, input := range inputs {
		inputInts := integers(testcases.ParseIntegerFile(input))
		inputInts.QuickSort()
		for i, v := range inputInts {
			assert.Equal(t, i+1, v)
		}
	}
}
