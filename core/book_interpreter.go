package core

import (
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
	Battle   BookConfig
	Ambience Ambience
	Event    Event
	Player   BookConfig
}

type BookConfig struct {
	Contents map[int]map[string][]string
	Type     string
}

func (book *Books) Load() {
	book.Lore = BookConfig(utils.Interpreter(LORE_BOOK))
	book.Battle = BookConfig(utils.Interpreter(BATTLE_BOOK))
	book.Player = BookConfig(utils.Interpreter(PLAYER_BOOK))

	ambience := BookConfig(utils.Interpreter(AMBIENCE_BOOK))
	book.Ambience = Ambience{Models: ambience.ToAmbience()}

	event := BookConfig(utils.Interpreter(EVENT_BOOK))
	book.Event = Event{Models: event.ToEvent()}
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

func (b BookConfig) ToEvent() []model.EventModel {
	event := []model.EventModel{}

	for _, item := range b.Contents {
		newEvent := model.EventModel{
			Contents:        item["#CONTENT"],
			System:          item["#SYSTEM"][0],
			EventTypeString: item["#EVENT_TYPE"][0],
			Type:            b.Type,
		}
		newEvent.PopulateEventType()

		event = append(event, newEvent)
	}

	return event
}
