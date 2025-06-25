package event

import "github.com/vOmarred11/VakuyProtocol/item"

// InteractEvent is when someone interact with a block or an item
type InteractEvent struct {
	// Item is the item that the client interacted with
	Item item.Item
}

func (e *InteractEvent) InteractEventItem() item.Item {
	return e.Item
}
