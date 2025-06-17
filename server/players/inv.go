package players

import (
	"VakuyProtocol/inv"
	"VakuyProtocol/stack"
)

// Inventory is the data of the player inventory
type Inventory struct {
	// Items is a slice of all items in the inventory
	Items []stack.Item
}

func (i Inventory) InventoryDataItems() []stack.Item {
	return i.Items
}
func (i Inventory) Len() int32 {
	m := inv.InventoryData{}
	return m.NonEmptySlots()
}
