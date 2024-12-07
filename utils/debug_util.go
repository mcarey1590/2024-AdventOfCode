package utils

import (
	"fmt"
	"os"
)

func WriteLinesToFile[T any](input []T, filename string) {
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
