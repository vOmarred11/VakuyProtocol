package pk

import (
	"github.com/vOmarred11/VakuyProtocol/player"
)

// VerticalCollision is when a player collides with another object or entity
type VerticalCollision struct {
	// Position is where player collided
	Position player.Position
}

func (h HorizontalCollision) VerticalCollisionPosition() player.Position {
	return h.Position
}
