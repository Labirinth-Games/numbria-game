package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	Numbria "github.com/Joeverson/numbria-game/src"
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

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
	}

	// init game

	//fmt.Print(logo) // pegar od arquivo na riaz
	// time.Sleep(3 * time.Second)

	storytelling := Numbria.StoryTelling{Current: Numbria.INITIAL_CHAPTER_INDEX, Book: books}
	storytelling.Execute()

	for {
		fmt.Print("\n[Você diz]: ")
		text, _ := reader.ReadString('\n')

		text = strings.ToLower(text)
		text = strings.Replace(text, "\n", "", -1)

		response, ok := Numbria.Brain(text, books.Dictionary)

		if ok {
			switch response.CommandType {
			case "Player":
				player.Invoke(&ctx, response.Action, text, response.Response)
			case "Event":
				ctx.Event.Invoke(&ctx, response.Action, text, response.Response)
			case "Battle":
				ctx.Battle.Invoke(&ctx, response.Action, text, response.Response)
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
	}
}
