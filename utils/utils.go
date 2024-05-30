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

func ChooseRandom(texts []string) string {
	return texts[rand.IntN(len(texts))]
}
