package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tolbier/algorithms_go/lib/testcases"

	"github.com/tolbier/algorithms_go/course1/week1/multiplier"
)

func Test_Multiplier_Multiply(t *testing.T) {
	inputs := [][]string{
		{"4", "2"},
		{"23", "1"},
		{"43", "2"},
		{"44", "21"},
		{"12323541", "21"},
		{"21", "12323541"},
		{"1234", "5678"},
		{"345", "890"},
		{"12345", "67890"},
		{"10920459779307117", "11281235394849568"},
		{"3141592653589793238462643383279502884197169399375105820974944592", "2718281828459045235360287471352662497757247093699959574966967627"},
	}
	results := []string{
		"8",
		"23",
		"86",
		"924",
		"258794361",
		"258794361",
		"7006652",
		"307050",
		"838102050",
		"123196277390350550270285486775456",
		"8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184",
	}
	for idx, result := range results {
		a := multiplier.Multiply(inputs[idx][0], inputs[idx][1])
		assert.Equal(t, result, a)

	}
}
func Test_Multiplier_Multiply_TestCases(t *testing.T) {
	inputs, outputs := testcases.InputsOutputs("testcases")
	for idx, input := range inputs {
		inputsTexts := testcases.ParseTextFile(input)
		outputsTexts := testcases.ParseTextFile(outputs[idx])

		result := multiplier.Multiply(inputsTexts[0], inputsTexts[1])
		assert.Equal(t, result, outputsTexts[0])
	}
}
