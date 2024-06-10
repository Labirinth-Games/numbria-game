package Numbria

import (
	"fmt"
	"strings"

	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

type Skill struct {
	Models []model.SkillModel
}

func NewSkill(paper utils.InterpreterConfig) Skill {

	skill := Skill{}

	for _, item := range paper.Book {
		skill.Models = append(skill.Models, model.SkillModel{
			Name:              utils.GetFirst("#NAME", item),
			Index:             utils.GetFirst("#INDEX", item),
			Description:       utils.GetFirst("#DESCRIPTION", item),
			EnergyPoint:       utils.GetFirstToInt("#ENERGY_POINT", item),
			RemainEnergyPoint: utils.GetFirstToInt("#ENERGY_POINT", item),
			Hanking:           types.HankingType(utils.GetFirst("#HANKING", item)),
		})
	}

	return skill
}

func (s *Skill) SerializerDictionary(interpreter utils.InterpreterConfig) []model.Dictionary {
	var dictionary []model.Dictionary

	for _, item := range interpreter.Book {
		var inputs []string

		for _, ask := range item["#NAME"] {
			words := strings.Split(ask, " ")

			for _, el := range words {
				if !utils.ExistsStringInArray(el, inputs) && len(el) > 2 {
					inputs = append(inputs, strings.ToLower(el))
				}
			}
		}

		dictionary = append(dictionary, model.Dictionary{
			Inputs:      inputs,
			Response:    []string{},
			CommandType: interpreter.Type,
			Action:      "Combate",
			Priority:    utils.GetFirstToInt("#PRIORITY", item),
			Index:       utils.GetFirst("#INDEX", item),
		})

	}

	return dictionary
}

func (s Skill) GetSkill(index string) *model.SkillModel {
	skillStorage := utils.Find(s.Models, func(current model.SkillModel) bool {
		return strings.Compare(index, current.Index) == 0
	})

	skill := skillStorage //clone

	return &skill
}

func (s Skill) PlayerSkillInfo(p Player) {
	var content []string

	fmt.Printf("\n-------------------------- Skills Player ---------------------------")
	fmt.Printf("\n\n\n\033[0;32m%s%s\033[0m", utils.TableItem("Name", 25), utils.TableItem("Remaing Energy Point", 25))

	for _, item := range p.Stats.Skills {
		name := utils.TableItem(item.Name, 25)
		point := utils.TableItem(fmt.Sprintf("%d", item.RemainEnergyPoint), 25)

		content = append(content, fmt.Sprintf("\n%s%s", name, point))
	}
	fmt.Printf("%s\n", strings.Join(content, ""))
	fmt.Printf("\n-------------------------------------------------------------------")
}
