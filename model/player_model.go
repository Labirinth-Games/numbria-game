package model

import (
	"github.com/Joeverson/numbria-game/utils"
)

const SKILL_BOOK = "./books/dictionaries/skills.dictionary"

type PlayerModel struct {
	Name     string
	HP       int
	Accuracy int
	Skills   []*SkillModel
	Strength utils.Dice
	Level    int
}

func (p *PlayerModel) Create(name string) {
	p.Name = name
	p.HP = 20
	p.Accuracy = 3
	p.Strength = utils.Dices.D8
}

func (p *PlayerModel) Hit(damage int) {

}

func (p *PlayerModel) Die() {

}
