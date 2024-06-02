package model

import (
	"log"

	"github.com/Joeverson/numbria-game/utils"
)

const (
	BESTIARY_BOOK = "./books/bestiary.book"
)

type CreatureModel struct {
	Name           string
	Stats          StatsModel
	Description    string
	ObserverSucess []string
	Level          int
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
	c.ObserverSucess = creature.ObserverSucess
	c.Stats = creature.Stats
}

func (c *CreatureModel) Hit(damage int) {

}

func (c *CreatureModel) Attack() int {
	if utils.TestPrecision(c.Stats.Accuracy) {

		return utils.RollDice(c.Stats.Strength)
	}

	return 0
}

func (c *CreatureModel) Die() {

}

func Serializer(data utils.InterpreterConfig) []CreatureModel {
	creature := []CreatureModel{}

	for _, item := range data.Book {
		if len(item) == 0 {
			continue
		}

		creature = append(creature, CreatureModel{
			Name:           utils.GetFirst("#NAME", item),
			Description:    utils.GetFirst("#DESCRIPTION", item),
			ObserverSucess: item["#OBSERVER_SUCESS"],
			Stats: StatsModel{
				HP:       utils.GetFirstToInt("#HP", item),
				Accuracy: utils.GetFirstToInt("#ACCURACY", item),
				Strength: utils.ConvertToDiceEnum(utils.GetFirst("#STRENGTH", item)),
			},
		})
	}

	return creature
}
