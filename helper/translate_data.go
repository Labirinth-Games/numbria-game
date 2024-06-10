package helper

import (
	"strings"

	"github.com/Joeverson/numbria-game/model"
)

func TranslateMultiplyTextToStorageData(texts []string, storage model.Storage) []string {
	var results []string

	for _, text := range texts {
		results = append(results, TranslateTextToStorageData(text, storage))
	}

	return results
}

func TranslateTextToStorageData(text string, storage model.Storage) string {
	if len(text) <= 0 {
		return ""
	}

	var newText []string

	split := strings.Split(text, " ")
	if len(split) == 0 {
		return ""
	}

	for _, word := range split {
		if len(word) > 3 && strings.Compare(word[:1], "<") == 0 && strings.Compare(word[len(word)-1:], ">") == 0 {
			key := word[1 : len(word)-1]

			word = storage.Get(key).(string)
		}

		newText = append(newText, word)
	}

	return strings.Join(newText, " ")
}
