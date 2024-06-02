package model

import "github.com/Joeverson/numbria-game/types"

type EventModel struct {
	Contents        []string
	System          string
	Fail            []string
	EventTypeEnum   types.EventType
	EventTypeString string
	Type            string
}

func (e *EventModel) PopulateEventType() {
	switch e.EventTypeString {
	case "creature":
		e.EventTypeEnum = types.EventTypeEnum.Creature
	case "explore":
		e.EventTypeEnum = types.EventTypeEnum.Explore
	}
}

func (e *EventModel) IsCreature() bool {
	return e.EventTypeEnum == types.EventTypeEnum.Creature
}
