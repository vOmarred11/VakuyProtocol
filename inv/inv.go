package inv

import (
	"github.com/vOmarred11/VakuyProtocol/stack"
)

const (
	InventoryTypeDefault = iota
	InventoryTypeCreative
	InventoryTypeSpectator
)

// InventoryData is data of the inventory
type InventoryData struct {
	// Items is a slice of all items
	Items []stack.Item
	// Slots are all available slots
	Slots int32
	// EmptySlots are all empty slots
	EmptySlots int32
	// Type is the type of the inventory
	Type uint8
}

func (d InventoryData) NonEmptySlots() int32 {
	v := d.EmptySlots - d.Slots
	return v
}
func (d InventoryData) InventoryDataItems() []stack.Item {
	return d.Items
}
func (d InventoryData) InventoryDataSlots() int32 {
	return d.Slots
}
func (d InventoryData) InventoryDataEmptySlot() int32 {
	return d.EmptySlots
}
func (d InventoryData) InventoryDataType() uint8 {
	return d.Type
}
