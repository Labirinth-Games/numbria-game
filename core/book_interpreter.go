package core

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/utils"
)

const (
	BATTLE_BOOK   = "./books/battle.book"
	LORE_BOOK     = "./books/lore.book"
	EVENT_BOOK    = "./books/event.book"
	AMBIENCE_BOOK = "./books/ambience.book"
	PLAYER_BOOK   = "./books/player.book"
)

type Books struct {
	Lore     BookConfig
	Event    BookConfig
	Battle   BookConfig
	Ambience Ambience
	Player   BookConfig
}

type BookConfig struct {
	Contents map[int]map[string][]string
	Type     string
}

func (book *Books) Load() {
	loreBook := ReadFileToStringList(LORE_BOOK)
	eventBook := ReadFileToStringList(EVENT_BOOK)
	battleBook := ReadFileToStringList(BATTLE_BOOK)
	ambienceBook := ReadFileToStringList(AMBIENCE_BOOK)
	playerBook := ReadFileToStringList(PLAYER_BOOK)

	book.Lore.Interpreter(loreBook)
	book.Event.Interpreter(eventBook)
	book.Battle.Interpreter(battleBook)
	book.Player.Interpreter(playerBook)

	ambience := BookConfig{}
	ambience.Interpreter(ambienceBook)
	book.Ambience = Ambience{Models: ambience.ToAmbience()}
}

func ReadFileToStringList(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func (b *BookConfig) Interpreter(lines []string) {
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

	b.Contents = interpretMap
	b.Type = bookType
}

func (b BookConfig) ToPlayer() []model.Dictionary {
	var dictionary []model.Dictionary

	for _, item := range b.Contents {
		var asks []string

		for _, ask := range item["#ASK"] {
			asklist := strings.Split(ask, " ")

			for _, el := range asklist {
				if !utils.ExistsStringInArray(el, asks) && len(el) > 2 {
					asks = append(asks, strings.ToLower(el))
				}
			}
		}

		var action string
		var priority int = 10

		if item["#ACTION"] != nil {
			action = item["#ACTION"][0]
		}

		if item["#PRIORITY"] != nil {
			value, err := strconv.Atoi(item["#PRIORITY"][0])

			if err != nil {
				panic(err)
			}

			priority = value
		}

		dictionary = append(dictionary, model.Dictionary{
			Inputs:   asks,
			Response: item["#ANSWER"],
			Type:     b.Type,
			Action:   action,
			Priority: priority,
		})
	}

	return dictionary
}

func (b BookConfig) ToAmbience() []model.AmbienceModel {
	ambience := []model.AmbienceModel{}

	for _, item := range b.Contents {
		ambience = append(ambience, model.AmbienceModel{Narration: item["#NARRATION"], Type: item["#TYPE"][0]})
	}

	return ambience
}
