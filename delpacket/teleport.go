package delpacket

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

type Teleport struct {
	// Enabled defines if teleport is enabled
	Enabled bool
	// TODO figure out what is this for
	EntityRuntimeID uint32
	// Position is the position where player will be teleported
	// You must insert 3 values so x, y, z
	Position [][][]mgl32.Vec3
	// Range is the teleport range
	Range float32
	// ChunkRadius is all chunks between the started chunk and final chunk
	ChunkRadius protocol.ChunkPos
}
type Position struct {
	X float32
	Y float32
	Z float32
}

func (x *Teleport) ID() uint32 {
	return x.ID()
}
func (x *Teleport) Marshal(r protocol.IO, p Position) {
	r.Bool(&x.Enabled)
	r.Vec3(&x.Position[int32(p.X)][int32(p.Y)][int32(p.Z)])
	r.Float32(&x.Range)
	r.ChunkPos(&x.ChunkRadius)
	r.Uint32(&x.EntityRuntimeID)
}
