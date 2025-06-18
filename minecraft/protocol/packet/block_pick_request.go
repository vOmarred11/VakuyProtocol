package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// BlockPickRequest is sent by the client when it requests to pick a block in the world and place its item in
// their inventory.
type BlockPickRequest struct {
	// Position is the position at which the client requested to pick the block. The block at that position
	// should have its item put in HotBarSlot if it is empty.
	Position protocol.BlockPos
	// AddBlockNBT specifies if the item should get all NBT tags from the block, meaning the item places a
	// block practically always equal to the one picked.
	AddBlockNBT bool
	// HotBarSlot is the slot that was held at the time of picking a block.
	HotBarSlot byte
}

// ID ...
func (*BlockPickRequest) ID() uint32 {
	return IDBlockPickRequest
}

func (pk *BlockPickRequest) Marshal(io protocol.IO) {
	io.BlockPos(&pk.Position)
	io.Bool(&pk.AddBlockNBT)
	io.Uint8(&pk.HotBarSlot)
}
