package pos

import (
	"github.com/vOmarred11/VakuyProtocol/event"
	"github.com/vOmarred11/VakuyProtocol/player"
)

// EventPosition returns the position on an event
type EventPosition struct {
	// Event is the event data
	Event event.Event
	// Position is the position on event
	Position player.Position
}

func (e EventPosition) EventPositionEvent() event.Event {
	return e.Event
}
func (e EventPosition) EventPosition() player.Position {
	return e.Position
}
