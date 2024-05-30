package core

import (
	"sort"
	"strings"

	"github.com/Joeverson/numbria-game/model"
)

type Match struct {
	Hash       string
	Rate       float32
	dictionary model.Dictionary
}

func Process(input string, dictionaries []model.Dictionary) model.Dictionary {
	var matchs []Match

	// quebrar o tesxto que recebeu
	inputTextArr := strings.Split(input, " ")

	// buscar em todos os dicionarios que foram treinados e apartir das palavras de entrada tenta achar a ação a ser tomada
	for _, dictionary := range dictionaries {
		amountOccurs := amountOccurrences(inputTextArr, dictionary.Inputs)

		if amountOccurs > 0 {
			matchs = append(matchs, Match{dictionary: dictionary, Rate: float32(amountOccurs) / float32(len(inputTextArr))})
		}
	}

	return win(matchs)
}

func amountOccurrences(originArr []string, targetArr []string) int {
	var occors int = 0

	for _, a := range originArr {
		for _, b := range targetArr {
			if strings.Compare(a, b) == 0 {
				occors++
				break
			}
		}
	}

	return occors
}

func win(matchs []Match) model.Dictionary {
	sort.Slice(matchs, func(i, j int) bool {
		return matchs[j].Rate < matchs[i].Rate
	})

	var concuring []Match

	for i, match := range matchs {
		if i+1 >= len(matchs) || match.Rate != matchs[i+1].Rate {
			break
		}

		if i == 0 {
			concuring = append(concuring, match)
		}

		concuring = append(concuring, matchs[i+1])
	}

	sort.Slice(concuring, func(i, j int) bool {
		return matchs[j].dictionary.Priority > matchs[i].dictionary.Priority
	})

	if len(matchs) > 0 {
		return matchs[0].dictionary
	}

	return model.Dictionary{}
}
