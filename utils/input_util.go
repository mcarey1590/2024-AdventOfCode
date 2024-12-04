package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(inputPath string) ([]string, error) {

	contents, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	lines := strings.Split(string(contents), "\n")
	return lines, nil
}
