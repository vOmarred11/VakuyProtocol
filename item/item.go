package item

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ItemStackMax = 64
	ItemStackMin = 1
)

type Item struct {
	// Name is the minecraft name of the item
	// Example "minecraft: dirt"
	Name string
	// CustomName is the custom name of the item, if this field is empty the result will be the minecraft item name
	CustomName string
	// NetworkID is the network id of the item
	NetworkID int32
	// Droppable defines if the item can be dropped
	Droppable bool
}

func (x Item) Hash(io protocol.IO) {
	io.String(&x.Name)
	io.String(&x.CustomName)
	io.Int32(&x.NetworkID)
	io.Bool(&x.Droppable)
}

type ItemStack struct {
	Item  Item
	Items []Item
}

func (x ItemStack) Hash() ItemStack {
	return x.Hash()
}
