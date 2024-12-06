package utils

import (
	"fmt"
	"os"
)

func WriteIntsToFile(input [][]int, filename string) {
	inputStrings := make([][]string, len(input))
	for i, line := range input {
		inputStrings[i] = make([]string, len(line))
		for j, num := range line {
			inputStrings[i][j] = fmt.Sprint(num)
		}
		WriteLinesToFile(inputStrings, filename)
	}
}

func WriteLinesToFile(input [][]string, filename string) {
	if filename == "" {
		filename = "output.txt"
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	for _, line := range input {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
