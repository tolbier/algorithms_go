package rselect

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/courses/course1/week4/rselect"
	"github.com/tolbier/algorithms_go/lib/testcases"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Test_RSelect_TestCases(t *testing.T) {
	inputs, _ := testcases.InputsOutputs("testcases")
	for _, input := range inputs {
		inputInts := rselect.Integers(testcases.ParseIntegerFile(input))
		slc := make(rselect.Integers, len(inputInts))
		for i := 0; i < 40; i++ {
			copy(slc, inputInts)
			j := rand.Intn(len(slc))
			v := slc.RSelect(j)
			assert.Equal(t, j+1, v)
		}
	}
}
