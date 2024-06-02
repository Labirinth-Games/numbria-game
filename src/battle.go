package Numbria

import (
	"fmt"
	"math/rand/v2"

	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

type Battle struct {
	Models   []model.BattleModel
	IsBattle bool
}

func NewBattle(paper utils.InterpreterConfig) Battle {
	battle := Battle{}

	for _, item := range paper.Book {
		newEvent := model.BattleModel{
			Narration: item["#NARRATION"],
			Status:    utils.GetFirst("#STATUS", item),
		}

		battle.Models = append(battle.Models, newEvent)
	}

	return battle
}

func RollIniciative(ctx *Context) {
	dice := utils.RollDice(utils.Dices.D6)
	var win string

	if dice >= 3 {
		ctx.Iniciative = types.IniciativeType.Player
		win = "sua"
	} else {
		ctx.Iniciative = types.IniciativeType.Enemy
		win = "do inimigo"
	}

	utils.SystemMultiplyDialog([]string{
		"Vamos rolar 1d6 para ver a iniciativa...",
		fmt.Sprintf("Você tirou %d, a iniciativa é %s", dice, win),
	}, 100)
}

func IniciativeToggle(ctx *Context) {
	if ctx.IsIniciativePlayer() {
		ctx.Iniciative = types.IniciativeType.Enemy
	} else {
		ctx.Iniciative = types.IniciativeType.Player
	}
}

func (b *Battle) GetNarrationSucess() string {
	for _, model := range b.Models {
		if model.Status == string(types.BattleStatusType.Sucess) {
			return utils.Random(model.Narration)
		}
	}

	return ""
}

func (b *Battle) GetNarrationFail() string {
	for _, model := range b.Models {
		if model.Status == string(types.BattleStatusType.Fail) {
			return utils.Random(model.Narration)
		}
	}

	return ""
}

/* -------------------------------------------------------------------------- */
/*                                Actions                                     */
/* -------------------------------------------------------------------------- */

// func (b *Battle) AttackCreature(target model.CreatureModel) {
// 	damage := creature.Attack()
// 	ctx.Player.Hit(damage)

// 	utils.NarrationDialog(ctx.Battle.GetNarrationSucess(), creature.Name)
// }

func (b *Battle) AttackPlayer(ctx *Context) {
	creature := ctx.Creatures[0]
	damage := creature.Attack()

	ctx.Player.Hit(damage)

	utils.NarrationDialog(b.GetNarrationFail(), creature.Name)
}

/* -------------------------------------------------------------------------- */
/*                                Utils Dynamic                               */
/* -------------------------------------------------------------------------- */

func (e *Battle) Invoke(ctx *Context, funcName string, args ...interface{}) {

	var ActionsMapper = map[string]interface{}{}

	action, ok := ActionsMapper[funcName]

	if !ok {
		if ctx.InBattle && ctx.InEvent {
			utils.SystemDialog("Não entendi o que voce quis fazer.")
			return
		}

		text := args[0].(string)
		answers := args[1].([]string)

		utils.NarrationDialog(answers[rand.IntN(len(answers))], text)
		return
	}

	action.(func(ctx *Context, text string, answers []string))(ctx, args[0].(string), args[1].([]string))
}
