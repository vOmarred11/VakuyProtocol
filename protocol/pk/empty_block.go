package pk

import "github.com/vOmarred11/VakuyProtocol/pos"

// EmptyBlock defines all empty blocks
// this case is not that common, it happens when the world is outdated
// for example when the world version is 1.16 but your version is 1.15
// some blocks will miss and minecraft doesn't count those missing block as air
type EmptyBlock struct {
	// Position defines the position of each block
	Position pos.BlockPos
	// Amount is the amount of blocks
	Amount int64
}

func (e EmptyBlock) EmptyBlockPosition() pos.BlockPos {
	return e.Position
}
func (e EmptyBlock) EmptyBlockAmount() int64 {
	return e.Amount
}
