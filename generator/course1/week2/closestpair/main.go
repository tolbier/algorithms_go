package main

import (
	"fmt"
	"os"

	"github.com/tolbier/algorithms_go/lib/testcases"
)

const testcasesDir = "course1/week2/closestpair/tests/testcases/"

func main() {
	generate_test(1, 2)
	generate_test(2, 4)
	generate_test(3, 8)
	generate_test(4, 16)
	generate_test(5, 32)
	generate_test(6, 64)
	generate_test(7, 128)
	generate_test(8, 256)
	generate_test(9, 512)
}

func generate_test(id int, n int) {
	inputFileName := fmt.Sprintf("%sinput_%d_%d.txt", testcasesDir, id, n)
	outputFileName := fmt.Sprintf("%soutput_%d_%d.txt", testcasesDir, id, n)

	testcases.WriteFile(inputFileName, func(file *os.File) (err error) {
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				fmt.Fprintf(file, "%6.2f %6.2f\n", float64(x), float64(y))
			}
		}
		fmt.Fprintln(file, "1 0.7")
		return
	})
	testcases.WriteFile(outputFileName, func(file *os.File) (err error) {
		fmt.Fprintln(file, "1 0.7")
		fmt.Fprintln(file, "1 1")
		return
	})
}
