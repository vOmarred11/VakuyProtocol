package delpacket

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	FallingModeNormal = iota
	FallingModeRotation
	// FallingModeTeleport cannot be called a mode because it teleports your instantly to the closest place
	// This is used when you are in a cave, and you want to get up
	FallingModeTeleport
	FallingModePhase
) // Top teleports your up
type Top struct {
	// TODO figure out what is this for
	EntityRuntimeID uint32
	// Value is the value of the x-asix
	Value mgl32.Vec3
	// FallingMode is the mode when your falling
	FallingMode uint8
}

func (x *Top) ID() uint32 { return x.ID() }

func (x *Top) Marshal(r protocol.IO) {
	r.Uint32(&x.EntityRuntimeID)
	r.Vec3(&x.Value)
	r.Uint8(&x.FallingMode)
}
