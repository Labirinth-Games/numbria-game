package Numbria

import (
	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/utils"
)

type Lore struct {
	Models []model.LoreModel
}

func NewLore(paper utils.InterpreterConfig) Lore {
	lore := Lore{}

	for _, item := range paper.Book {
		lore.Models = append(lore.Models, model.LoreModel{
			Content:       item["#CONTENT"],
			Read:          utils.GetFirst("#READ", item),
			Save:          utils.GetFirst("#SAVE", item),
			Next:          utils.GetFirst("#NEXT", item),
			NextAutomatic: utils.Exist("#NEXT_AUTOMATIC", item),
			Index:         utils.GetFirst("#INDEX", item),
			System:        utils.GetFirst("#SYSTEM", item),
		})
	}

	return lore
}

func (l Lore) FindByIndex(index string) (model.LoreModel, bool) {
	for _, item := range l.Models {
		if index == item.Index {
			return item, true
		}
	}

	return model.LoreModel{}, false
}
