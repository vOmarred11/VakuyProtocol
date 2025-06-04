package delpacket

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Velocity struct {
	// if Enabled is true reach is active
	Enabled bool
	// TODO figure out what is this for
	EntityRuntimeID uint64
	// Value is the value of the velocity
	Value float32
}

func (x *Velocity) ID() uint32 {
	return x.ID()
}
func (x Velocity) Marshal(r protocol.IO) {
	r.Bool(&x.Enabled)
	r.Uint64(&x.EntityRuntimeID)
	r.Float32(&x.Value)
}
