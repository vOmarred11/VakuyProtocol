package event

import "github.com/vOmarred11/VakuyProtocol/stack"

// InteractEvent is when someone interact with a block or an item
type InteractEvent struct {
	// Item is the item that the client interacted with
	Item stack.Item
}

func (e *InteractEvent) InteractEventItem() stack.Item {
	return e.Item
}
