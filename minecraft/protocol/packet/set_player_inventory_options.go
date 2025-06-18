package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	InventoryLayoutNone = iota
	InventoryLayoutSurvival
	InventoryLayoutRecipeBook
	InventoryLayoutCreative
)

const (
	InventoryLeftTabNone = iota
	InventoryLeftTabConstruction
	InventoryLeftTabEquipment
	InventoryLeftTabItems
	InventoryLeftTabNature
	InventoryLeftTabSearch
	InventoryLeftTabSurvival
)

const (
	InventoryRightTabNone = iota
	InventoryRightTabFullScreen
	InventoryRightTabCrafting
	InventoryRightTabArmour
)

// SetPlayerInventoryOptions is a bidirectional packet that can be used to update the inventory options of a player.
type SetPlayerInventoryOptions struct {
	// LeftInventoryTab is the tab that is selected on the left side of the inventory. This is usually for the creative
	// inventory. It is one of the InventoryLeftTab constants above.
	LeftInventoryTab byte
	// RightInventoryTab is the tab that is selected on the right side of the inventory. This is usually for the player's
	// own inventory. It is one of the InventoryRightTab constants above.
	RightInventoryTab byte
	// Filtering is whether the player has enabled the filtering between recipes they have unlocked or not.
	Filtering bool
	// InventoryLayout is the layout of the inventory. It is one of the InventoryLayout constants above.
	InventoryLayout byte
	// CraftingLayout is the layout of the crafting inventory. It is one of the InventoryLayout constants above.
	CraftingLayout byte
}

// ID ...
func (*SetPlayerInventoryOptions) ID() uint32 {
	return IDSetPlayerInventoryOptions
}

func (pk *SetPlayerInventoryOptions) Marshal(io protocol.IO) {
	io.Uint8(&pk.LeftInventoryTab)
	io.Uint8(&pk.RightInventoryTab)
	io.Bool(&pk.Filtering)
	io.Uint8(&pk.InventoryLayout)
	io.Uint8(&pk.CraftingLayout)
}
