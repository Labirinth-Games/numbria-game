package Numbria

import (
	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/types"
)

type Context struct {
	Creatures    []*model.CreatureModel
	CurrentEvent model.EventModel
	Ambience     Ambience
	Event        Event
	Battle       Battle
	Player       Player
	Skill        Skill
	Storage      *model.Storage

	Iniciative types.Iniciative

	InBattle     bool
	InEvent      bool
	IsTypingSave bool
}

func (c *Context) IsIniciativePlayer() bool {
	return c.Iniciative == types.IniciativeType.Player
}

func (c *Context) IsIniciativeEnemy() bool {
	return c.Iniciative == types.IniciativeType.Enemy
}
