package utils

import (
	"bufio"
	"os"
	"strings"
)

type InterpreterConfig struct {
	Book map[int]map[string][]string
	Type string
}

func readBook(path string) ([]string, bool) {
	file, err := os.Open(path)

	if err != nil {
		return nil, false
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, true
}

func Interpreter(path string) InterpreterConfig {
	lines, ok := readBook(path)

	if !ok {
		panic("PATH FILE NOT FOUND " + strings.ToUpper(path))
	}

	interpretMap := make(map[int]map[string][]string)
	var currentLabel int
	var subCurrentLabel string
	var bookType string
	var i int = 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "__TYPE__:") {
			bookType = line[len("__TYPE__:"):]

			continue
		}

		if strings.Contains(line, "----------") {
			currentLabel = i
			i++

			interpretMap[currentLabel] = map[string][]string{}
			continue
		}

		if len(line) > 1 && line[:1] == "#" {
			subCurrentLabel = line
			interpretMap[currentLabel][subCurrentLabel] = []string{}
			continue
		}

		interpretMap[currentLabel][subCurrentLabel] = append(interpretMap[currentLabel][subCurrentLabel], line)
	}

	return InterpreterConfig{
		Book: interpretMap,
		Type: bookType,
	}
}
