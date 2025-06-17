package event

import (
	"VakuyProtocol/player"
)

const (
	MoveModeNormal = iota
	MoveModeTeleport
	MoveModeCollapse
)

// MoveEvent is when a player moves
type MoveEvent struct {
	// Position is the position of the player
	Position player.Position
	// Pitch is the vertical rotation of the player. Facing straight forward yields a pitch of 0. Pitch is
	// measured in degrees.
	Pitch float32
	// Yaw is the horizontal rotation of the player. Yaw is also measured in degrees.
	Yaw float32
	// Speed is the speed which the player does things
	Speed float32
	// Tick is the tick sent from the server
	Tick uint32
	// OnGround defines if the player is on the ground or no
	OnGround bool
	// Mode is the moving mode
	Mode uint8
}

func (m MoveEvent) MoveEventPosition() player.Position {
	return m.Position
}

func (m MoveEvent) MoveEventYaw() float32 {
	return m.Yaw
}

func (m MoveEvent) MoveEventPitch() float32 {
	return m.Pitch
}

func (m MoveEvent) MoveEventSpeed() float32 {
	return m.Speed
}

func (m MoveEvent) MoveEventTick() uint32 {
	return m.Tick
}

func (m MoveEvent) MoveEventOnGround() bool {
	return m.OnGround
}

func (m MoveEvent) MoveEventMode() uint8 {
	return m.Mode
}
