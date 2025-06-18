package players

import (
	"github.com/vOmarred11/VakuyProtocol/pos"
)

// Move is the defines target movement
type Move struct {
	// Farway is the distance sent by the server
	Farway float32
	// Distance is the distance received by the client and the target
	Distance float32
	// Detachment is the distance detected by the client when the target is on a combo-air only when a hit
	// gets detected
	Detachment float32
	// DetachmentAir is the distance detected by the target when the client is on a combo-air only when a hit
	// gets detected
	DetachmentAir float32
	// Yaw is the target yaw
	Yaw float32
	// Pitch is the target pitch
	Pitch float32
	// OnGround defines is the target is on air without hit
	OnGround bool
	// Position is the position of the target
	Position pos.Entity
}

func (p Move) PlayerFarway() float32 {
	return p.Farway
}
func (p Move) PlayerDistance() float32 {
	return p.Distance
}
func (p Move) PlayerDetachment() float32 {
	return p.Detachment
}
func (p Move) PlayerDetachmentAir() float32 {
	return p.DetachmentAir
}
func (p Move) PlayerYaw() float32 {
	return p.Yaw
}
func (p Move) PlayerPitch() float32 {
	return p.Pitch
}
func (p Move) PlayerOnGround() bool {
	return p.OnGround
}
func (p Move) PlayerPosition() pos.Entity {
	return p.Position
}
