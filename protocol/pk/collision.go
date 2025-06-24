package pk

import "github.com/vOmarred11/VakuyProtocol/pos"

const (
	CollisionTypeNormal = iota
	CollisionTypeTeleport
	CollisionTypeRotation
)

// Collision defines any type of collision
type Collision struct {
	// Position is the position where player collided
	Position pos.BlockPos
}

func (c Collision) CollisionPosition() pos.BlockPos {
	return c.Position
}
