package model

import (
	"github.com/Joeverson/numbria-game/types"
)

type SkillModel struct {
	Index             string
	Name              string
	Description       string
	Hanking           types.HankingType
	EnergyPoint       int
	RemainEnergyPoint int
}
