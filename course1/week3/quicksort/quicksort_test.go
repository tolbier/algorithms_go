package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/lib/testcases"
)

func Test_QuickSort_TestCases(t *testing.T) {
	inputs, outputs := testcases.InputsOutputs("testcases")
	for idx, input := range inputs {
		inputInts := testcases.ParseIntegerFile(input)
		outputInts := testcases.ParseIntegerFile(outputs[idx])
		result := ThreeQuickSorts(inputInts)
		assert.Equal(t, result, outputInts)

	}
}

func ThreeQuickSorts(inputInts []int) (res []int) {
	slc := make(integers, len(inputInts))
	copy(slc, inputInts)
	res = append(res, slc.QuickSort(NewFirstPivoter()))
	copy(slc, inputInts)
	res = append(res, slc.QuickSort(NewLastPivoter()))
	copy(slc, inputInts)
	res = append(res, slc.QuickSort(NewMedianOfThreePivoter()))
	return
}
