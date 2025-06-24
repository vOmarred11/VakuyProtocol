package pk

import "github.com/vOmarred11/VakuyProtocol/player"

// AirWalk is when a player is walking on the air
type AirWalk struct {
	// Height is height where player is walking
	// usually this is always the same value because if it will change
	// it will be detected by AirJump
	Height player.Position
	// Speed is the speed that the player is walking only on the air
	Speed float32
}

func (a AirWalk) AirWalkHeight() player.Position {
	return a.Height
}
func (a AirWalk) AirWalkSpeed() float32 {
	return a.Speed
}
