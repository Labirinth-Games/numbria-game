package controller

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/Joeverson/numbria-game/core"
	"github.com/Joeverson/numbria-game/game"
	"github.com/Joeverson/numbria-game/utils"
)

const (
	DELAY_TO_SHOW_MESSAGE = 10 // milliseconds 1500 default
	INITIAL_CHAPTER_INDEX = "001"
)

type StoryTelling struct {
	Current int
	Book    core.Books
	Ctx     *game.Context
}

func (st StoryTelling) Execute() {
	lore, ok := st.Book.Lore.Contents[st.Current]

	if !ok {
		utils.SystemSay("Não encontrei o chapter")
	}

	for key, value := range lore {
		if strings.Compare(key, "#GOTO") == 0 {
			st.Goto(value[0])
			continue
		}

		if strings.Compare(key, "#SYSTEM") == 0 {
			utils.SystemSay(value[0])
			continue
		}

		if strings.Compare(key, "#CONTENT") == 0 {
			for _, text := range value {
				utils.UniverseSay(text)
				time.Sleep(DELAY_TO_SHOW_MESSAGE * time.Millisecond)
			}
		}

	}
}

func (st *StoryTelling) Next() error {
	currentIndex, ok := getCurrentIndexChapter(INITIAL_CHAPTER_INDEX, st.Book.Lore.Contents)

	if !ok {
		log.Fatal(errors.New("CHAPTER NOT FOUND"))
	}

	if len(st.Book.Lore.Contents) > currentIndex+1 {
		st.Current += 1

		st.Execute()

		return nil
	}

	return errors.New("CHAPTER IS FINISHED")
}

func (st *StoryTelling) Goto(index string) {
	utils.SpaceBlank()

	chapterIndex, ok := getCurrentIndexChapter(index, st.Book.Lore.Contents)

	if !ok {
		log.Fatal(errors.New("CHAPTER NOT FOUND"))
	}

	st.Current = chapterIndex
	st.Execute()
}

/* -------------------------------------------------------------------------- */
/*                                Utils                                       */
/* -------------------------------------------------------------------------- */

func getCurrentIndexChapter(index string, arr map[int]map[string][]string) (int, bool) {
	for i, chapter := range arr {
		for key, value := range chapter {
			if strings.Compare(key, "#INDEX") == 0 && value[0] == index {
				return i, true
			}
		}
	}

	return -1, false
}
