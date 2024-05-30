package utils

import "hash/fnv"

func Hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))

	return string(rune(h.Sum32()))
}
