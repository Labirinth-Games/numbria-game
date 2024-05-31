package types

type EventType int

var EventTypeEnum eventTypeEnum

const (
	creature EventType = iota + 1
	explore
)

type eventTypeEnum struct {
	Creature EventType
	Explore  EventType
}

func init() {
	EventTypeEnum = eventTypeEnum{
		Creature: creature,
		Explore:  explore,
	}
}
