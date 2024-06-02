package Numbria

import (
	"math/rand/v2"

	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

const (
	PROBABILITY_SPOT_EVENT = 1 // 0.3 // 30% de chance de ter evento
)

type Event struct {
	Models []model.EventModel
}

func NewEvent(paper utils.InterpreterConfig) Event {
	event := Event{}

	for _, item := range paper.Book {
		newEvent := model.EventModel{
			Contents:        item["#CONTENT"],
			Fail:            item["#FAIL"],
			System:          utils.GetFirst("#SYSTEM", item),
			EventTypeString: utils.GetFirst("#EVENT_TYPE", item),
			Type:            paper.Type,
		}
		newEvent.PopulateEventType()

		event.Models = append(event.Models, newEvent)
	}

	return event
}

func (e *Event) TryTriggerEvent() (event model.EventModel, element interface{}, ok bool) {
	if utils.IsProbable(PROBABILITY_SPOT_EVENT) {
		event = utils.Random(e.Models)

		switch event.EventTypeEnum {
		case types.EventTypeEnum.Creature:
			creature := model.CreatureModel{}
			creature.Create()

			element = creature
		}

		utils.NarrationMultiplyDialog(event.Contents, DELAY_TO_SHOW_MESSAGE)
		utils.SystemDialog(event.System)

		ok = true
		return
	}

	ok = false
	return
}

/* -------------------------------------------------------------------------- */
/*                                Actions                                     */
/* -------------------------------------------------------------------------- */

func (e *Event) observerAction(ctx *Context, text string, answers []string) {
	RollIniciative(ctx)

	if len(ctx.Creatures) == 0 {
		return
	}

	creature := ctx.Creatures[0]

	if ctx.IsIniciativePlayer() {
		if ctx.CurrentEvent.IsCreature() {
			utils.NarrationDialog(utils.Random(creature.ObserverSucess))
		}
	}

	if ctx.IsIniciativeEnemy() {
		utils.NarrationDialog(utils.Random(ctx.CurrentEvent.Fail))

		if ctx.CurrentEvent.IsCreature() {

			// NOTE - Quando o jogador tenta observar mas n consegue o monstro dem 40% de chande de atacar
			if utils.IsProbable(0.4) {
				ctx.Battle.AttackPlayer(ctx)
			}
		}
	}

	// utils.SystemDialog(ctx.CurrentEvent.System)
	ctx.InEvent = false
}

/* -------------------------------------------------------------------------- */
/*                                Utils Dynamic                               */
/* -------------------------------------------------------------------------- */

func (e *Event) Invoke(ctx *Context, funcName string, args ...interface{}) {

	if !ctx.InEvent {
		utils.SystemDialog(utils.Random([]string{
			"Não entendi o que voce quis fazer.",
			"Acho que voce endoidou, falando coisa com coisa",
			"Acho que não lhe entendi, o que vc quer?",
		}))
		return
	}

	var ActionsMapper = map[string]interface{}{
		"ObserverAction": e.observerAction,
		// "StartEventAttack":   e.StartEventAttack,
	}

	action, ok := ActionsMapper[funcName]

	if !ok {
		text := args[0].(string)
		answers := args[1].([]string)

		utils.NarrationDialog(answers[rand.IntN(len(answers))], text)
		return
	}

	action.(func(ctx *Context, text string, answers []string))(ctx, args[0].(string), args[1].([]string))
}
