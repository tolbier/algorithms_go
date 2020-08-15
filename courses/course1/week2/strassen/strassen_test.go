package strassen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Strassen_Multiply(t *testing.T) {
	a := Matrix{
		{1, 2},
		{5, 6},
	}.Multiply(
		Matrix{
			{3, 4},
			{7, 8},
		})
	assert.Equal(t, Matrix{
		{17, 20},
		{57, 68}}, a)
}
func Test_Strassen_Multiply2(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}.Multiply(
		Matrix{
			{-1, -2, -3, -4},
			{-5, -6, -7, -8},
			{-9, -10, -11, -12},
			{-13, -14, -15, -16},
		})
	assert.Equal(t, Matrix{
		{-90, -100, -110, -120},
		{-202, -228, -254, -280},
		{-314, -356, -398, -440},
		{-426, -484, -542, -600},
	}, a)
}
func Test_Strassen_Sum(t *testing.T) {
	b := Matrix{
		{1, 2},
		{5, 6},
	}
	c := Matrix{
		{3, 4},
		{7, 8},
	}
	a := b.Sum(c)
	assert.Equal(t, Matrix{
		{4, 6},
		{12, 14}}, a)
}
func Test_Strassen_Substract(t *testing.T) {
	a := Matrix{
		{1, 2},
		{5, 6},
	}.Subtract(
		Matrix{
			{3, 4},
			{7, 8},
		})
	assert.Equal(t, Matrix{
		{-2, -2},
		{-2, -2}}, a)
}
