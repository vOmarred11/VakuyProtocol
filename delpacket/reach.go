package delpacket

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Reach struct {
	// if Enabled is true reach is active
	Enabled bool
	// EntityRuntimeID is the runtime id for the entity
	EntityRuntimeID uint64
	// Distance is the block distance
	Distance float32
}

func (x *Reach) ID() uint32 {
	return x.ID()
}
func (x Reach) Marshal(r protocol.IO) {
	r.Bool(&x.Enabled)
	r.Uint64(&x.EntityRuntimeID)
	r.Float32(&x.Distance)
}
