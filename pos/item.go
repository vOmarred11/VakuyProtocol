package pos

import "github.com/vOmarred11/VakuyProtocol/item"

// ItemPosition returns the position of a chosen item when is not in the inventory
type ItemPosition struct {
	// Item is the chosen item
	Item item.ItemStack
	X    float64
	Y    float64
	Z    float64
}

func (i ItemPosition) ItemPositionItem() item.ItemStack {
	return i.Item
}
func (i ItemPosition) ItemPositionX() float64 {
	return i.X
}
func (i ItemPosition) ItemPositionY() float64 {
	return i.Y
}
func (i ItemPosition) ItemPositionZ() float64 {
	return i.Z
}
