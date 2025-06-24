package pk

import "github.com/vOmarred11/VakuyProtocol/player"

// SprintCollision is when a player collides while sprinting
type SprintCollision struct {
	// Position is where player collided
	Position player.Position
}

func (s SprintCollision) SprintCollisionPosition() player.Position {
	return s.Position
}
