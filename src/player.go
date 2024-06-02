package Numbria

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/utils"
)

type Player struct {
	World World
	Stats *model.PlayerModel

	x         int
	y         int
	direction string
}

func (p *Player) Load(world World) {
	stats := &model.PlayerModel{}
	stats.Create("PlayerName")

	p.Stats = stats
	p.World = world
}

func (player *Player) Spawn() {
	player.x = rand.IntN(player.World.Width-2) + 2
	player.y = rand.IntN(player.World.Height-2) + 2
}

func (p Player) getPlaceName() string {
	return p.World.GetNameZone(p.x, p.y)
}

func (p *Player) move() {
	if strings.Contains(p.direction, "norte") {
		p.y += 1
	}

	if strings.Contains(p.direction, "sul") {
		p.y -= 1
	}

	if strings.Contains(p.direction, "leste") {
		p.x += 1
	}

	if strings.Contains(p.direction, "oeste") {
		p.x -= 1
	}
}

func (p *Player) Hit(damage int) {
	p.Stats.HP -= damage

	utils.SystemDialog(fmt.Sprintf("Voce sofreu %d de dano", damage))

	if p.Stats.HP <= 0 {
		p.Die()
	}
}

func (p *Player) Die() {
	utils.SystemDialog("\t\t ================= YOU DIE ================ \n\n")
	os.Exit(0)
}

func (p *Player) Attack() int {
	if utils.TestPrecision(p.Stats.Accuracy) {
		return utils.RollDice(p.Stats.Strength)
	}

	return 0
}

/* -------------------------------------------------------------------------- */
/*                                Actions                                     */
/* -------------------------------------------------------------------------- */

func (p Player) WhereIAm(ctx *Context, text string, answers []string) {
	utils.NarrationDialog(utils.Random(answers), p.getPlaceName())
}

func (p Player) WhatsThere(ctx *Context, text string, answers []string) {
	var direction string

	for _, dir := range []string{"norte", "sul", "leste", "oeste"} {
		if strings.Contains(dir, text) {
			direction = dir
			break
		}
	}

	message := fmt.Sprintf(utils.Random(answers), direction, p.getPlaceName())

	utils.NarrationDialog(message)
}

func (p *Player) Walk(ctx *Context, text string, answers []string) {
	p.direction = utils.ExtractString(text, []string{"norte", "sul", "leste", "oeste"})

	p.move()

	hasEvent := eventProcess(ctx)

	if hasEvent {
		return
	}

	ctx.Ambience.TalkAbout(p.getPlaceName())
	utils.NarrationDialog(utils.Random(answers), p.direction)
}

/* -------------------------------------------------------------------------- */
/*                                Utils                                       */
/* -------------------------------------------------------------------------- */

func eventProcess(ctx *Context) bool {
	if ctx.InEvent {
		return false
	}

	event, subEvent, hasEvent := ctx.Event.TryTriggerEvent()

	if hasEvent {
		ctx.InEvent = true
		if event.IsCreature() {
			creature := subEvent.(model.CreatureModel)

			ctx.Creatures = append(ctx.Creatures, creature)
			ctx.CurrentEvent = event
		}
		return true
	}

	return false
}

/* -------------------------------------------------------------------------- */
/*                                Commands                                    */
/* -------------------------------------------------------------------------- */

func (p Player) GetPositionInfo() {
	fmt.Printf("\n x:%d y:%d - place: %s\n\n", p.x, p.y, p.getPlaceName())
}

func (p Player) StatsInfo() {
	fmt.Printf("\n Hp: %d, Accuracy: %d \n\n", p.Stats.HP, p.Stats.Accuracy)
}

/* -------------------------------------------------------------------------- */
/*                                Utils Dynamic                               */
/* -------------------------------------------------------------------------- */

func (p *Player) Invoke(ctx *Context, funcName string, args ...interface{}) {

	var ActionsMapper = map[string]interface{}{
		"Walk":       p.Walk,
		"WhereIAm":   p.WhereIAm,
		"WhatsThere": p.WhatsThere,
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
