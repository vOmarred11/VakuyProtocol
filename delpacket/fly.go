package delpacket

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	FlyModeCreative = iota
	FlyModeNoClip
)

// Fly is a Packet that allows you to fly
type Fly struct {
	// TODO Figure out what is this for
	EntityRuntimeID uint32
	// Mode is the mode that is used
	Mode byte
	// if Enabled is set to true you can fly
	Enabled bool
}

func (x *Fly) ID() uint32 {
	return x.ID()
}

func (x *Fly) Marshal(r protocol.IO) {
	r.Uint32(&x.EntityRuntimeID)
	r.Bool(&x.Enabled)
	r.Uint8(&x.Mode)
}
