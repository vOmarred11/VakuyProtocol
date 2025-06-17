package event

import (
	"VakuyProtocol/player"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// BreakBlockEvent is when a player breaks a block
type BreakBlockEvent struct {
	// Block is the block id
	Block int32
	// Position is the position of the block
	Position protocol.BlockPos
	// BreakingSpeed is the speed that the block will break
	BreakingSpeed float32
}
type BreakBlockData struct {
	Player player.Player
}

func (e BreakBlockEvent) BreakBlockEventBlock() int32 {
	return e.Block
}
func (e BreakBlockEvent) BreakBlockEventPosition() protocol.BlockPos {
	return e.Position
}
func (d BreakBlockData) PlayerData() player.Player {
	return d.Player
}
