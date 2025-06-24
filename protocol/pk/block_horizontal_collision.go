package pk

import "github.com/vOmarred11/VakuyProtocol/pos"

// BlockHorizontalCollision is when a player collides with a block horizontally
type BlockHorizontalCollision struct {
	// Type is the type of the collision
	Type uint8
	// Position is where the player collided
	Position pos.BlockPos
}

func (b BlockHorizontalCollision) BlockHorizontalCollisionType() uint8 {
	return b.Type
}
func (b BlockHorizontalCollision) BlockHorizontalCollisionPosition() pos.BlockPos {
	return b.Position
}
