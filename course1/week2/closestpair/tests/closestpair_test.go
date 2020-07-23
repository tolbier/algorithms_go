package tests

import (
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/course1/week2/closestpair"
	"github.com/tolbier/algorithms_go/course1/week2/closestpair/model"
	"github.com/tolbier/algorithms_go/lib/testcases"
)

func Test_ClosestPair_TestCases(t *testing.T) {
	inputs, outputs := testcases.InputsOutputs("testcases")
	for idx, input := range inputs {
		inputPoints := ParsePointsFile(input)
		outputPoints := ParsePointsFile(outputs[idx])
		p := closestpair.Grid(inputPoints).ClosestPair() //return a PointPair
		assert.True(t,
			p.Is(*outputPoints[0], *outputPoints[1]))
	}
}
func ParsePointsFile(filename string) (res []*model.Point) {
	textLines := testcases.ParseTextFile(filename)
	r := regexp.MustCompile(`\W*(\d*\.?\d*)\W*(\d*\.?\d*)$`)
	for _, line := range textLines {
		sm := r.FindStringSubmatch(line)
		x, _ := strconv.ParseFloat(sm[1], 64)
		y, _ := strconv.ParseFloat(sm[2], 64)
		p := &model.Point{X: x, Y: y}
		res = append(res, p)
	}
	return res
}
