package pk

import (
	"github.com/vOmarred11/VakuyProtocol/player"
)

// HorizontalCollision is when a player collides with another object or entity
type HorizontalCollision struct {
	// Position is where player collided
	Position player.Position
}

func (h HorizontalCollision) HorizontalCollisionPosition() player.Position {
	return h.Position
}
