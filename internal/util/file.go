package util

import (
	"bufio"
	"os"
)

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}

func ReadLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
