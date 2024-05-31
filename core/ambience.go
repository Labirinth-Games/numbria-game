package core

import (
	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/utils"
)

type Ambience struct {
	Models []model.AmbienceModel
}

func (a Ambience) TalkAbout(ambienceType string) {
	find, ok := a.find(ambienceType)

	if ok {
		utils.UniverseSay(utils.Random(find.Narration))
	}
}

func (a Ambience) find(ambienceType string) (model.AmbienceModel, bool) {
	for _, model := range a.Models {
		if model.Type == ambienceType {
			return model, true
		}
	}

	return model.AmbienceModel{}, false
}
