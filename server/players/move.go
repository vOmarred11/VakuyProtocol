package players

import (
	"VakuyProtocol/pos"
)

// Player is the nearest player to the client
type Player struct {
	// Farway is the distance sent by the server
	Farway float32
	// Distance is the distance received by the client and the player
	Distance float32
	// Detachment is the distance detected by the client when the player is on a combo-air only when a hit
	// gets detected
	Detachment float32
	// DetachmentAir is the distance detected by the player when the client is on a combo-air only when a hit
	// gets detected
	DetachmentAir float32
	// Yaw is the player yaw
	Yaw float32
	// Pitch is the player pitch
	Pitch float32
	// OnGround defines is the player is on air without hit
	OnGround bool
	// Position is the position of the player
	Position pos.Entity
}

func (p Player) PlayerFarway() float32 {
	return p.Farway
}
func (p Player) PlayerDistance() float32 {
	return p.Distance
}
func (p Player) PlayerDetachment() float32 {
	return p.Detachment
}
func (p Player) PlayerDetachmentAir() float32 {
	return p.DetachmentAir
}
func (p Player) PlayerYaw() float32 {
	return p.Yaw
}
func (p Player) PlayerPitch() float32 {
	return p.Pitch
}
func (p Player) PlayerOnGround() bool {
	return p.OnGround
}
func (p Player) PlayerPosition() pos.Entity {
	return p.Position
}
