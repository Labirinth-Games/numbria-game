package utils

import (
	"math/rand/v2"
	"strconv"
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

func Exist(key string, mapper map[string][]string) bool {
	_, ok := mapper[key]

	return ok
}

func GetFirst(key string, mapper map[string][]string) string {
	value, ok := mapper[key]

	if ok {
		return value[0]
	}

	return ""
}

func GetFirstToInt(key string, mapper map[string][]string) int {
	value, ok := mapper[key]

	if !ok {
		return 0
	}

	intValue, err := strconv.Atoi(value[0])

	if err != nil {
		panic(err)
	}

	return intValue
}

func IsProbable(percent float32) bool {
	return rand.Float32() <= percent
}

func Find[T any](arr []T, cb func(current T) bool) (result T) {
	for _, item := range arr {
		if cb(item) {
			result = item
			return
		}
	}

	return
}
