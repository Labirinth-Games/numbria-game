package utils

import (
	"math/rand/v2"
	"strings"
)

func ExistsStringInArray(val string, arr []string) bool {
	for _, el := range arr {
		if strings.Compare(val, el) == 0 {
			return true
		}
	}

	return false
}

func ExtractString(input string, find []string) string {
	arr := strings.Split(input, " ")

	for _, el := range arr {
		for _, item := range find {
			if strings.Compare(item, el) == 0 {
				return item
			}
		}
	}

	return ""
}

func Random[T any](elements []T) T {
	return elements[rand.IntN(len(elements))]
}
