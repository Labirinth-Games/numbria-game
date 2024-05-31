package core

import (
	"math/rand/v2"

	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

const (
	PROBABILITY_SPOT_EVENT = 0.3 // 30% de chance de ter evento
)

type Event struct {
	Models []model.EventModel
}

func (e Event) TriggerEvent() (event model.EventModel, element interface{}, isOk bool) {
	if rand.Float32() <= PROBABILITY_SPOT_EVENT {
		event = utils.Random(e.Models)

		switch event.EventTypeEnum {
		case types.EventTypeEnum.Creature:
			element = createCreature()
		}

		utils.DelayDialog(event.Contents, 10)
		utils.SystemSay(event.System)

		isOk = true
		return
	}

	isOk = false
	return
}

func createCreature() model.CreatureModel {
	creature := model.CreatureModel{}
	creature.Create()

	return creature
}
