package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Joeverson/numbria-game/model"
	Numbria "github.com/Joeverson/numbria-game/src"
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

var Context Numbria.Context

func main() {
	// loaders

	reader := bufio.NewReader(os.Stdin)

	world := Numbria.World{}
	world.MapGenerate()

	books := Numbria.Book{}
	books.Load()

	player := Numbria.Player{}
	player.Load(world)
	player.Spawn()

	ctx := Numbria.Context{
		Iniciative: types.IniciativeType.None,
		Event:      books.Event,
		Ambience:   books.Ambience,
		Battle:     books.Battle,
		Player:     player,
		Skill:      books.Skill,
		Storage:    &model.StorageData,
	}

	Context = ctx

	player.ReceiveSkill(&ctx, "001")
	player.ReceiveSkill(&ctx, "002")

	// init game
	logo, ok := utils.ReadFileToString("./logo.txt")
	if ok {
		fmt.Print(logo) // pegar od arquivo na riaz
		time.Sleep(3 * time.Second)
	}

	storytelling := Numbria.StoryTelling{Current: Numbria.INITIAL_CHAPTER_INDEX, Book: books}

	for {
		lore := storytelling.GetLore(&ctx)
		storytelling.Play(ctx)

		utils.DisplayCommandIndicator(ctx.InBattle)

		text, _ := reader.ReadString('\n')
		text = strings.ToLower(text)
		text = strings.Replace(text, "\n", "", -1)

		if ctx.IsTypingSave {
			ctx.Storage.Save(lore.Save, text)
			ctx.IsTypingSave = false
			continue
		}

		response, ok := Numbria.Brain(text, books.Dictionary)

		if ok {
			switch response.CommandType {
			case "player":
				player.Invoke(&ctx, response.Action, text, response.Response)
			case "event":
				ctx.Event.Invoke(&ctx, response.Action, text, response.Response)
			case "battle":
				ctx.Battle.Invoke(&ctx, response.Action, response)
			}
		}

		/* -------------------------------------------------------------------------- */
		/*                                System                                      */
		/* -------------------------------------------------------------------------- */

		if utils.ExistsStringInArray(text, []string{"exit", "!e"}) {
			fmt.Println("\nsaindo...")
			os.Exit(0)
		}

		if utils.ExistsStringInArray(text, []string{"!coor", "!c", "!coord"}) {
			player.GetPositionInfo()
		}

		if utils.ExistsStringInArray(text, []string{"!stats", "!s"}) {
			player.StatsInfo()
		}

		if utils.ExistsStringInArray(text, []string{"!skill", "!sk"}) {
			ctx.Skill.PlayerSkillInfo(player)
		}
	}
}
