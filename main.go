package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Joeverson/numbria-game/controller"
	"github.com/Joeverson/numbria-game/core"
	"github.com/Joeverson/numbria-game/game"
	"github.com/Joeverson/numbria-game/persona"
	"github.com/Joeverson/numbria-game/utils"
)

const logo = `
----------------------//\\
--------------------// ¤ \\
--------------------\\ ¤ //
--------------------- \\//
-------------------- (___)
---------------------(___)
---------------------(___)
---------------------(___)_________
----------\_____/\__/----\__/\_____/
------------\ _°_[-викинг-]_ _° /
----------------\_°_¤ ---- ¤_°_/
--------------------\ __°__ /
---------------------|\_°_/|
---------------------[|\_/|]
---------------------[|[¤]|]
---------------------[|;¤;|]       ::      ::  ::    ::  ::          :: :: ::::    ::::::    ::   ::::::
---------------------[;;¤;;]       ::::    ::  ::    ::  ::::      :::: ::     ::  ::    ::  ::  ::    ::  
--------------------;;;;¤]|]\      ::  ::  ::  ::    ::  ::  ::  ::  :: :: ::::::  :: ::     ::  :: :: ::   
-------------------;;;;;¤]|]-\     ::    ::::  ::    ::  ::    ::    :: ::     ::  ::   ::   ::  ::    ::   
------------------;;;;;[¤]|]--\    ::      ::  :: :: ::  ::          :: :: ::::    ::    ::  ::  ::    ::  
-----------------;;;;;|[¤]|]---\
----------------;;;;;[|[¤]|]|---| 
---------------;;;;;[|[¤]|]|---|
----------------;;;;[|[¤]|/---/
-----------------;;;[|[¤]/---/
------------------;;[|[¤/---/
-------------------;[|[/---/
--------------------[|/---/
---------------------/---/
--------------------/---/|]
-------------------/---/]|];
------------------/---/¤]|];;
-----------------|---|[¤]|];;;
-----------------|---|[¤]|];;;
------------------\--|[¤]|];;
-------------------\-|[¤]|];
---------------------\|[¤]|]
----------------------\\¤//
-----------------------\|/
------------------------V

`

func main() {
	// loaders

	reader := bufio.NewReader(os.Stdin)

	world := game.World{}
	world.MapGenerate()

	books := core.Books{}
	books.Load()

	player := persona.Player{}
	player.Load(world, books)
	player.Spawn()

	ctx := game.Context{}

	// init game

	fmt.Print(logo)
	// time.Sleep(3 * time.Second)

	storytelling := controller.StoryTelling{Current: 0, Book: books}

	storytelling.Execute()

	for {
		fmt.Print("\n[Você diz]: ")
		text, _ := reader.ReadString('\n')

		text = strings.ToLower(text)
		text = strings.Replace(text, "\n", "", -1)

		response := core.Process(text, player.Dictionary)

		if response.Type != "" {
			switch response.Type {
			case "Player":
				player.Invoke(&ctx, response.Action, text, response.Response)
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
	}
}
