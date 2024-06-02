package Numbria

import (
	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/utils"
)

type Ambience struct {
	Models []model.AmbienceModel
}

func NewAmbience(paper utils.InterpreterConfig) Ambience {
	ambience := Ambience{}

	for _, item := range paper.Book {
		ambience.Models = append(ambience.Models, model.AmbienceModel{
			Narration: item["#NARRATION"],
			Type:      utils.GetFirst("#TYPE", item),
		})
	}

	return ambience
}

func (a *Ambience) TalkAbout(ambienceType string) {
	find, ok := a.find(ambienceType)

	if ok {
		utils.NarrationDialog(utils.Random(find.Narration))
	}
}

func (a *Ambience) find(ambienceType string) (model.AmbienceModel, bool) {
	for _, model := range a.Models {
		if model.Type == ambienceType {
			return model, true
		}
	}

	return model.AmbienceModel{}, false
}
