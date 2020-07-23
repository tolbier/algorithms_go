package testcases

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func ParseIntegerFile(filename string) (res []int) {
	textLines := ParseTextFile(filename)
	for _, line := range textLines {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func WriteFile(filename string, f func(*os.File) error) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = f(file)
	return err
}
func ParseTextFile(filename string) (res []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(scanner.Text())
		log.Fatal(err)
	}
	return res
}

func InputsOutputs(root string) (inputs []string, outputs []string) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			output, err := outputFile(path)
			if err == nil {
				inputs = append(inputs, path)
				outputs = append(outputs, output)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return inputs, outputs
}

func outputFile(path string) (string, error) {
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	r := regexp.MustCompile(`^input_(.+)$`)
	submatch := r.FindStringSubmatch(base)
	if len(submatch) == 2 {
		return fmt.Sprintf("%s/output_%s", dir, submatch[1]), nil
	}
	return "", errors.New("no match input file")
}
