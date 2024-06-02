package model

import (
	"fmt"
	"log"

	"github.com/Joeverson/numbria-game/utils"
)

const (
	BESTIARY_BOOK = "./books/bestiary.book"
)

type CreatureModel struct {
	Name                    string
	Stats                   *StatsModel
	Description             string
	NarrationObserverSucess []string
	NarrationAttackSucess   []string
	NarrationAttackFail     []string
	Level                   int
}

type StatsModel struct {
	HP       int
	Accuracy int
	Strength utils.Dice
}

func (c *CreatureModel) Create() {
	creatureBook := utils.Interpreter(BESTIARY_BOOK)
	creatures := Serializer(creatureBook)

	if len(creatures) == 0 {
		log.Fatalln("Error when get creatures on book")
		return
	}

	creature := utils.Random(creatures)

	c.Name = creature.Name
	c.Description = creature.Description
	c.NarrationObserverSucess = creature.NarrationObserverSucess
	c.NarrationAttackFail = creature.NarrationAttackFail
	c.NarrationAttackSucess = creature.NarrationAttackSucess

	c.Stats = creature.Stats
}

func (c *CreatureModel) Hit(damage int) {
	c.Stats.HP -= damage

	utils.SystemDialog(fmt.Sprintf("%s recebeu %d de dano", c.Name, damage))
}

func (c *CreatureModel) Attack() int {
	if utils.TestPrecision(c.Stats.Accuracy) {
		return utils.RollDice(c.Stats.Strength)
	}

	return 0
}

func (c *CreatureModel) IsDie() bool {
	return c.Stats.HP <= 0
}

func Serializer(data utils.InterpreterConfig) []CreatureModel {
	creature := []CreatureModel{}

	for _, item := range data.Book {
		if len(item) == 0 {
			continue
		}

		creature = append(creature, CreatureModel{
			Name:                    utils.GetFirst("#NAME", item),
			Description:             utils.GetFirst("#DESCRIPTION", item),
			NarrationObserverSucess: item["#OBSERVER_SUCESS"],
			NarrationAttackSucess:   item["#NARRATION_ATTACK_SUCESS"],
			NarrationAttackFail:     item["#NARRATION_ATTACK_FAIL"],
			Stats: &StatsModel{
				HP:       utils.GetFirstToInt("#HP", item),
				Accuracy: utils.GetFirstToInt("#ACCURACY", item),
				Strength: utils.ConvertToDiceEnum(utils.GetFirst("#STRENGTH", item)),
			},
		})
	}

	return creature
}
