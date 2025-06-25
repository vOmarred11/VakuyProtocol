package nbt

import "github.com/vOmarred11/VakuyProtocol/inventory"

// NBT is the nbt data item/block
type NBT struct {
	// Data represent the objects data
	Data []byte
	// Amount is the objects amount
	Amount int32
	// Inventory is the inventory data of the nbt
	Inventory inventory.InventoryData
}

func (n NBT) NBTData() []byte {
	return n.Data
}
func (n NBT) NBTAmount() int32 {
	return n.Amount
}
func (n NBT) NBTInventory() inventory.InventoryData {
	return n.Inventory
}
