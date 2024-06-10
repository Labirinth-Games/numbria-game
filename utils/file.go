package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileToString(path string) (string, bool) {
	file, err := os.Open(path)

	if err != nil {
		return "", false
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return strings.Join(lines, "\n"), true
}
