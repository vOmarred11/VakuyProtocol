package event

import (
	"VakuyProtocol/player"
	"VakuyProtocol/stack"
)

// DropItemEvent is when a player drops an item
type DropItemEvent struct {
	// Position is the position where the item got dropped
	// If Position is mid-block it will the result will be unknown
	Position stack.ItemPosition
	// Item is the item that got dropped
	Item stack.Item
}
type DropItemData struct {
	Player player.Player
}

func (d DropItemEvent) DropItemEventPosition() stack.ItemPosition {
	return d.Position
}
func (d DropItemEvent) DropItemEventItem() stack.Item {
	return d.Item
}
func (d DropItemData) DropItemEventPlayer() player.Player {
	return d.Player
}
