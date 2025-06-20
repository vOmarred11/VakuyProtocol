package nbt

import "github.com/vOmarred11/VakuyProtocol/inv"

// NBT is the nbt data item/block
type NBT struct {
	// Data represent the objects data
	Data []byte
	// Amount is the objects amount
	Amount int32
	// Inventory is the inventory data of the nbt
	Inventory inv.InventoryData
}

func (n NBT) NBTData() []byte {
	return n.Data
}
func (n NBT) NBTAmount() int32 {
	return n.Amount
}
func (n NBT) NBTInventory() inv.InventoryData {
	return n.Inventory
}
