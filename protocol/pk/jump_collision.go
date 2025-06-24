package pk

import "github.com/vOmarred11/VakuyProtocol/player"

// JumpCollision is when a player collides with a jump
type JumpCollision struct {
	// Position is where player collided
	Position player.Position
}
