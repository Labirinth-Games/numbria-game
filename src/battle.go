package Numbria

import (
	"fmt"

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

func (b *Battle) PlayerNarrationSucess() string {
	for _, model := range b.Models {
		if model.Status == string(types.BattleStatusType.Sucess) {
			return utils.Random(model.Narration)
		}
	}

	return ""
}

func (b *Battle) PlayerNarrationFail() string {
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

func (b *Battle) EnemyAttack(ctx *Context, skill *model.SkillModel) bool {
	creature := ctx.Creatures[0]
	damage := creature.Attack()

	if damage == 0 {
		utils.NarrationDialog(utils.Random(creature.NarrationAttackFail))
		return true
	}

	utils.NarrationDialog(utils.Random(creature.NarrationAttackSucess))
	ctx.Player.Hit(damage)

	return !ctx.Player.IsDie()
}

func (b *Battle) PlayerAttack(ctx *Context, skill *model.SkillModel) bool {
	if skill.RemainEnergyPoint <= 0 {
		utils.SystemDialog("Você está cansado demais para usar essa habilidade.")
		return true
	}

	damage := ctx.Player.Attack(skill)
	creature := ctx.Creatures[0]

	skill.RemainEnergyPoint -= 1

	if damage == 0 {
		utils.NarrationDialog(fmt.Sprintf(b.PlayerNarrationFail(), creature.Name))
		return true
	}

	utils.NarrationDialog(fmt.Sprintf(b.PlayerNarrationSucess(), creature.Name))
	creature.Hit(damage)

	return !creature.IsDie()
}

func (b *Battle) Combat(ctx *Context, skill *model.SkillModel) {
	if !ctx.InBattle {
		RollIniciative(ctx)
		ctx.InBattle = true
	}

	var first, second func(*Context, *model.SkillModel) bool = b.PlayerAttack, b.EnemyAttack

	if ctx.IsIniciativeEnemy() {
		first, second = second, first
	}

	if first(ctx, skill) {
		second(ctx, skill)
	}

	if ctx.Creatures[0].IsDie() {
		//TODO - drop item no futuro

		utils.DisplaySession("Fim da Batalha")

		ctx.InBattle = false
		ctx.InEvent = false

	}
}

/* -------------------------------------------------------------------------- */
/*                                Utils Dynamic                               */
/* -------------------------------------------------------------------------- */

func (e *Battle) Invoke(ctx *Context, funcName string, dictinonary model.Dictionary) {
	if !ctx.InEvent {
		utils.SystemDialog(utils.Random([]string{
			"Não entendi o que voce quis fazer.",
			"Acho que voce endoidou, falando coisa com coisa",
			"Acho que não lhe entendi, o que vc quer?",
		}))
		return
	}

	if dictinonary.Index == "" {
		utils.SystemDialog("Não entendi que habiliade é essa")
		return
	}

	skill, ok := ctx.Player.HasSkill(dictinonary.Index)

	if !ok {
		utils.SystemDialog(utils.Random([]string{
			"Você não conhece essa habilidade, voce precisa se dedicar mais nos estudos para poder usar-lá",
			"Você não sabe fazer isso, é invendando coisas aleatorias que as pessoas morrem...",
			"Que habilidade é essa? você não conhece",
		}))
		return
	}

	var ActionsMapper = map[string]interface{}{
		"Combate": e.Combat,
	}

	action, ok := ActionsMapper[funcName]

	if !ok {
		return
	}

	action.(func(ctx *Context, skill *model.SkillModel))(ctx, skill)
}
