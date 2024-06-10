package Numbria

import (
	"strings"

	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/utils"
)

const (
	BATTLE_BOOK   = "./books/battle.book"
	LORE_BOOK     = "./books/lore.book"
	EVENT_BOOK    = "./books/event.book"
	AMBIENCE_BOOK = "./books/ambience.book"

	PLAYER_DICTIONARY = "./books/player.dictionary"
	EVENT_DICTIONARY  = "./books/event.dictionary"
	BATTLE_DICTIONARY = "./books/battle.dictionary"
	SKILL_DICTIONARY  = "./books/dictionaries/skills.dictionary"
)

type Book struct {
	Lore       Lore
	Battle     Battle
	Ambience   Ambience
	Event      Event
	Skill      Skill
	Dictionary []model.Dictionary
}

func (book *Book) Load() {
	book.Lore = NewLore(utils.Interpreter(LORE_BOOK))
	book.Battle = NewBattle(utils.Interpreter(BATTLE_BOOK))
	book.Ambience = NewAmbience(utils.Interpreter(AMBIENCE_BOOK))
	book.Event = NewEvent(utils.Interpreter(EVENT_BOOK))
	book.Skill = NewSkill(utils.Interpreter(SKILL_DICTIONARY))

	book.Dictionary = SerializerDictionary([]utils.InterpreterConfig{
		utils.Interpreter(PLAYER_DICTIONARY),
		utils.Interpreter(EVENT_DICTIONARY),
		utils.Interpreter(BATTLE_DICTIONARY),
	})

	book.Dictionary = append(book.Dictionary, book.Skill.SerializerDictionary(utils.Interpreter(SKILL_DICTIONARY))...)
}

func SerializerDictionary(interpreters []utils.InterpreterConfig) []model.Dictionary {
	var dictionary []model.Dictionary

	for _, interpreter := range interpreters {
		for _, item := range interpreter.Book {
			var asks []string

			for _, ask := range item["#ASK"] {
				asklist := strings.Split(ask, " ")

				for _, el := range asklist {
					if !utils.ExistsStringInArray(el, asks) && len(el) > 2 {
						asks = append(asks, strings.ToLower(el))
					}
				}
			}

			dictionary = append(dictionary, model.Dictionary{
				Inputs:      asks,
				Response:    item["#ANSWER"],
				CommandType: interpreter.Type,
				Action:      utils.GetFirst("#ACTION", item),
				Priority:    utils.GetFirstToInt("#PRIORITY", item),
			})
		}
	}

	return dictionary
}
