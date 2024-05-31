package persona

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/Joeverson/numbria-game/core"
	"github.com/Joeverson/numbria-game/game"
	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

type Player struct {
	World      game.World
	book       core.Books
	Dictionary []model.Dictionary
	x          int
	y          int
	direction  string
}

func (p *Player) Load(world game.World, book core.Books) {

	p.Dictionary = book.Player.ToPlayer()
	p.World = world
	p.book = book
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

/* -------------------------------------------------------------------------- */
/*                                Actions                                     */
/* -------------------------------------------------------------------------- */

func (p Player) WhereIAm(ctx *game.Context, text string, answers []string) {
	utils.UniverseSay(utils.Random(answers), p.getPlaceName())
}

func (p Player) WhatsThere(ctx *game.Context, text string, answers []string) {
	var direction string

	for _, dir := range []string{"norte", "sul", "leste", "oeste"} {
		if strings.Contains(dir, text) {
			direction = dir
			break
		}
	}

	message := fmt.Sprintf(utils.Random(answers), direction, p.getPlaceName())

	utils.UniverseSay(message)
}

func (p *Player) Walk(ctx *game.Context, text string, answers []string) {
	p.direction = utils.ExtractString(text, []string{"norte", "sul", "leste", "oeste"})

	p.move()

	hasEvent := eventProcess(ctx, p)

	if hasEvent {
		return
	}

	p.book.Ambience.TalkAbout(p.getPlaceName())
	utils.UniverseSay(utils.Random(answers), p.direction)
}

func eventProcess(ctx *game.Context, p *Player) bool {
	event, subEvent, hasEvent := p.book.Event.TriggerEvent()

	if hasEvent {
		if event.EventTypeEnum == types.EventTypeEnum.Creature {
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

/* -------------------------------------------------------------------------- */
/*                                Utils Dynamic                               */
/* -------------------------------------------------------------------------- */

func (p *Player) Invoke(ctx *game.Context, name string, args ...interface{}) {
	var ActionsMapper = map[string]interface{}{
		"Walk":       p.Walk,
		"WhereIAm":   p.WhereIAm,
		"WhatsThere": p.WhatsThere,
	}

	action, ok := ActionsMapper[name]

	if !ok {
		text := args[0].(string)
		answers := args[1].([]string)

		utils.UniverseSay(answers[rand.IntN(len(answers))], text)
		return
	}

	action.(func(ctx *game.Context, text string, answers []string))(ctx, args[0].(string), args[1].([]string))
}
