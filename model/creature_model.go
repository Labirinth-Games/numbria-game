package model

import (
	"log"
	"strconv"

	"github.com/Joeverson/numbria-game/utils"
)

const (
	BESTIARY_BOOK = "./books/bestiary.book"
)

type CreatureModel struct {
	Name        string
	Stats       StatsModel
	Description string
	Level       int
}

type StatsModel struct {
	HP       int
	Accuracy int
	Strength utils.Dice
}

func (m *CreatureModel) Create() {
	creatureBook := utils.Interpreter(BESTIARY_BOOK)
	creatures := Serializer(creatureBook)

	if len(creatures) == 0 {
		log.Fatalln("Error when get creatures on book")
		return
	}

	creature := utils.Random(creatures)

	m.Name = creature.Name
	m.Description = creature.Description
	m.Stats = creature.Stats
}

func Serializer(data utils.InterpreterConfig) []CreatureModel {
	creature := []CreatureModel{}

	for _, item := range data.Contents {
		if len(item) == 0 {
			continue
		}

		HP, _ := strconv.Atoi(item["#HP"][0])
		accuracy, _ := strconv.Atoi(item["#ACCURACY"][0])

		creature = append(creature, CreatureModel{
			Name:        item["#NAME"][0],
			Description: item["#DESCRIPTION"][0],
			Stats: StatsModel{
				HP:       HP,
				Accuracy: accuracy,
				Strength: utils.ConvertToDiceEnum(item["#STRENGTH"][0]),
			},
		})
	}

	return creature
}
