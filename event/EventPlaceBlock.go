package event

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/vOmarred11/VakuyProtocol/player"
)

// PlaceBlockEvent is when a player breaks a block
type PlaceBlockEvent struct {
	// Block is the block id
	Block int32
	// Position is the position of the block
	Position protocol.BlockPos
}
type PlaceBlockData struct {
	Player player.Player
}

func (e PlaceBlockEvent) BreakBlockEventBlock() int32 {
	return e.Block
}
func (e PlaceBlockEvent) BreakBlockEventPosition() protocol.BlockPos {
	return e.Position
}
func (d PlaceBlockData) PlayerData() player.Player {
	return d.Player
}
