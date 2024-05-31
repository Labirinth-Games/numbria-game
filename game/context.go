package game

import (
	"github.com/Joeverson/numbria-game/model"
)

type Context struct {
	Context      interface{}
	Creatures    []model.CreatureModel
	CurrentEvent model.EventModel

	InBattle bool
}
