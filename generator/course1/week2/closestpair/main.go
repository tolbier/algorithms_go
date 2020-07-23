package main

import (
	"fmt"
	"os"

	"github.com/tolbier/algorithms_go/lib/testcases"
)

func main() {
	test1()
	test2()
}

func test1() {
	testcases.WriteFile("course1/week2/closestpair/tests/testcases/input_1_9.txt", func(file *os.File) (err error) {
		fmt.Fprintln(file, "1 1")
		fmt.Fprintln(file, "1 3")
		fmt.Fprintln(file, "1 5")
		fmt.Fprintln(file, "1 7")
		fmt.Fprintln(file, "2 1")
		fmt.Fprintln(file, "2 3")
		fmt.Fprintln(file, "2 5")
		fmt.Fprintln(file, "2 7")
		fmt.Fprintln(file, "1 2.7")
		return
	})
	testcases.WriteFile("course1/week2/closestpair/tests/testcases/output_1_9.txt", func(file *os.File) (err error) {
		fmt.Fprintln(file, "1 2.7")
		fmt.Fprintln(file, "1 3")
		return
	})
}
func test2() {
	testcases.WriteFile("course1/week2/closestpair/tests/testcases/input_2_9.txt", func(file *os.File) (err error) {
		for x := 0; x < 1000; x++ {
			for y := 0; y < 1000; y++ {
				fmt.Fprintf(file, "%6.2f %6.2f\n", float64(x), float64(y))

			}

		}
		fmt.Fprintln(file, "123 122.7")
		return
	})
	testcases.WriteFile("course1/week2/closestpair/tests/testcases/output_2_9.txt", func(file *os.File) (err error) {
		fmt.Fprintln(file, "123 122.7")
		fmt.Fprintln(file, "123 123")
		return
	})
}
