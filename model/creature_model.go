package model

import (
	"fmt"
	"log"
	"strings"

	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

const (
	BESTIARY_BOOK        = "./books/bestiary.book"
	BESTIARY_SKILLS_BOOK = "./books/bestiary_skills.book"
)

type CreatureModel struct {
	Name                    string
	Stats                   *StatsModel
	Description             string
	NarrationObserverSucess []string
	NarrationAttackSucess   []string
	NarrationAttackFail     []string
	NarrationDie            []string
	Skills                  []*SkillModel
	Level                   int
}

type StatsModel struct {
	HP       int
	Accuracy int
	Strength utils.Dice
}

func NewCreature() *CreatureModel {
	creatureBook := utils.Interpreter(BESTIARY_BOOK)
	creatureSkillsBook := utils.Interpreter(BESTIARY_SKILLS_BOOK)
	creatures := Serializer(creatureBook, creatureSkillsBook)

	if len(creatures) == 0 {
		log.Fatalln("Error when get creatures on book")
		return nil
	}

	creature := utils.Random(creatures)
	return &creature
}

func (c *CreatureModel) Hit(damage int) {
	c.Stats.HP -= damage

	utils.SystemDialog(fmt.Sprintf("%s recebeu %d de dano", c.Name, damage))

	if c.IsDie() {
		utils.NarrationDialog(utils.Random(c.NarrationDie))
	}
}

func (c *CreatureModel) Attack() int {
	if utils.TestPrecision(c.Stats.Accuracy) {
		skill := utils.Random(c.Skills)
		hankLevel := types.HankingTypeEnum.ToInt(skill.Hanking)

		return utils.RollDice(c.Stats.Strength) + hankLevel
	}

	return 0
}

func (c *CreatureModel) IsDie() bool {
	return c.Stats.HP <= 0
}

func Serializer(data utils.InterpreterConfig, skillsInterpreter utils.InterpreterConfig) []CreatureModel {
	creatures := []CreatureModel{}
	skills := skillsEnemySerializer(skillsInterpreter)

	for _, item := range data.Book {
		if len(item) == 0 {
			continue
		}

		skillIndexs, ok := item["#SKILLS"]
		var creatureSkills []*SkillModel = []*SkillModel{}

		if ok {
			creatureSkills = findByIndexArray(skillIndexs, skills)
		}

		creatures = append(creatures, CreatureModel{
			Name:                    utils.GetFirst("#NAME", item),
			Description:             utils.GetFirst("#DESCRIPTION", item),
			NarrationObserverSucess: item["#OBSERVER_SUCESS"],
			NarrationAttackSucess:   item["#NARRATION_ATTACK_SUCESS"],
			NarrationAttackFail:     item["#NARRATION_ATTACK_FAIL"],
			NarrationDie:            item["#NARRATION_DIE"],
			Skills:                  creatureSkills,
			Stats: &StatsModel{
				HP:       utils.GetFirstToInt("#HP", item),
				Accuracy: utils.GetFirstToInt("#ACCURACY", item),
				Strength: utils.ConvertToDiceEnum(utils.GetFirst("#STRENGTH", item)),
			},
		})
	}

	return creatures
}

func skillsEnemySerializer(data utils.InterpreterConfig) []*SkillModel {
	skills := []*SkillModel{}

	for _, item := range data.Book {
		if len(item) == 0 {
			continue
		}

		skills = append(skills, &SkillModel{
			Index:       utils.GetFirst("#INDEX", item),
			Name:        utils.GetFirst("#NAME", item),
			Description: utils.GetFirst("#DESCRIPTION", item),
			Hanking:     types.HankingType(utils.GetFirst("#HANKING", item)),
			EnergyPoint: utils.GetFirstToInt("#ENERGY_POINT", item),
		})
	}

	return skills
}

func findByIndexArray(indexs []string, arrSkills []*SkillModel) []*SkillModel {
	finded := []*SkillModel{}

	for _, index := range indexs {
		for _, item := range arrSkills {
			if strings.Compare(index, item.Index) == 0 {
				finded = append(finded, item)
			}
		}
	}

	return finded
}
