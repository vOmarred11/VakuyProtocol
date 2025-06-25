package event

import (
	"github.com/vOmarred11/VakuyProtocol/item"
	"github.com/vOmarred11/VakuyProtocol/player"
)

// DropItemEvent is when a player drops an item
type DropItemEvent struct {
	// Position is the position where the item got dropped
	// If Position is mid-block it will the result will be unknown
	Position item.ItemPosition
	// Item is the item that got dropped
	Item item.Item
}
type DropItemData struct {
	Player player.Player
}

func (d DropItemEvent) DropItemEventPosition() item.ItemPosition {
	return d.Position
}
func (d DropItemEvent) DropItemEventItem() item.Item {
	return d.Item
}
func (d DropItemData) DropItemEventPlayer() player.Player {
	return d.Player
}
