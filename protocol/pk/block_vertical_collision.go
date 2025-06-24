package pk

import "github.com/vOmarred11/VakuyProtocol/pos"

// BlockVerticalCollision is when a player collides with a block vertically
type BlockVerticalCollision struct {
	// Type is the type of the collision
	Type uint8
	// Position is where the player collided
	Position pos.BlockPos
}

func (b BlockHorizontalCollision) BlockVerticalCollisionType() uint8 {
	return b.Type
}
func (b BlockHorizontalCollision) BlockVerticalCollisionPosition() pos.BlockPos {
	return b.Position
}
