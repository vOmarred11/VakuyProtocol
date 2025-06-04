package event

import (
	"VakuyProtocol/event"
	"VakuyProtocol/player"
)

type Event struct {
	Player        player.Player
	Type          event.Event
	HasPermission bool
}

func (e *Event) EventPlayer() player.Player {
	return e.Player
}
func (e *Event) EventType() event.Event {
	return e.Type
}
func (e *Event) EventHasPermission() bool {
	if e.HasPermission == false {
		e.Player.Message("You do not have permission to perform this action")
	}
	return true
}
