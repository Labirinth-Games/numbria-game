package Numbria

import (
	"errors"
	"log"

	"github.com/Joeverson/numbria-game/utils"
)

const (
	DELAY_TO_SHOW_MESSAGE = 10 // milliseconds 1500 default
	INITIAL_CHAPTER_INDEX = "001"
)

type StoryTelling struct {
	Current string
	Book    Book
	Ctx     *Context
}

func (st StoryTelling) Execute() {
	lore, ok := st.Book.Lore.FindByIndex(st.Current)

	if !ok {
		utils.SystemDialog("Não encontrei o chapter")
	}

	if len(lore.Content) > 0 {
		utils.NarrationMultiplyDialog(lore.Content, DELAY_TO_SHOW_MESSAGE)
	}

	if lore.NextAutomatic {
		st.Next()
	}

	if lore.System != "" {
		utils.SystemDialog(lore.System)
	}
}

func (st *StoryTelling) Next() error {
	currentLore, ok := st.Book.Lore.FindByIndex(st.Current)

	if !ok {
		log.Fatal(errors.New("CHAPTER NOT FOUND"))
	}

	st.Current = currentLore.Next
	st.Execute()

	return nil
}

func (st *StoryTelling) Goto(index string) {
	utils.SpaceBlank()

	lore, ok := st.Book.Lore.FindByIndex(index)

	if !ok {
		log.Fatal(errors.New("CHAPTER NOT FOUND"))
	}

	st.Current = lore.Index
	st.Execute()
}
