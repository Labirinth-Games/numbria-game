package Numbria

import (
	"errors"
	"log"

	"github.com/Joeverson/numbria-game/helper"
	"github.com/Joeverson/numbria-game/model"
	"github.com/Joeverson/numbria-game/utils"
)

const (
	DELAY_TO_SHOW_MESSAGE = 500 // milliseconds 1500 default
	INITIAL_CHAPTER_INDEX = "001"
)

type StoryTelling struct {
	Current string
	Book    Book
}

func (st StoryTelling) GetCurrent() (model.LoreModel, bool) {
	return st.Book.Lore.FindByIndex(st.Current)
}

func (st StoryTelling) GetLore(ctx *Context) model.LoreModel {
	lore, ok := st.Book.Lore.FindByIndex(st.Current)

	if ok && lore.Save != "" {
		ctx.IsTypingSave = true
	}

	return lore
}

func (st *StoryTelling) Play(ctx Context) {
	lore, ok := st.GetCurrent()

	if !ok {
		utils.SystemDialog("Não encontrei o chapter")
	}

	if len(lore.Content) > 0 {
		utils.NarrationMultiplyDialog(helper.TranslateMultiplyTextToStorageData(lore.Content, *ctx.Storage), DELAY_TO_SHOW_MESSAGE)
	}

	if lore.NextAutomatic {
		st.AutomaticNext(ctx)
	}

	if lore.System != "" {
		utils.SystemDialog(helper.TranslateTextToStorageData(lore.System, *ctx.Storage))
	}

	st.Current = lore.Next
}

func (st *StoryTelling) AutomaticNext(ctx Context) error {
	currentLore, ok := st.GetCurrent()

	if !ok {
		log.Fatal(errors.New("CHAPTER NOT FOUND"))
	}

	st.Current = currentLore.Next
	st.Play(ctx)

	return nil
}

func (st *StoryTelling) Goto(index string, ctx Context) {
	utils.SpaceBlank()

	lore, ok := st.Book.Lore.FindByIndex(index)

	if !ok {
		log.Fatal(errors.New("CHAPTER NOT FOUND"))
	}

	st.Current = lore.Index
	st.Play(ctx)
}
