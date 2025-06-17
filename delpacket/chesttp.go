package delpacket

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ChestTp automatically teleports you to a chest
type ChestTp struct {
	// Enabled defines if ChestTp is enabled
	Enabled bool
	// TODO figure out what is this for
	EntityRuntimeID uint32
	// TeleportRange is the teleport range
	TeleportRange float32
	// ChestRange is the chest-search range
	ChestRange float32
	// ChunkRadius is all chunks between the started chunk and final chunk
	ChunkRadius protocol.ChunkPos
	// DoubleChests defines if you can teleport to double chests
	DoubleChests bool
}

func (x *ChestTp) ID() uint32 {
	return x.ID()
}
func (x *ChestTp) Marshal(r protocol.IO) {
	r.Uint32(&x.EntityRuntimeID)
	r.Float32(&x.TeleportRange)
	r.Float32(&x.ChestRange)
	r.Bool(&x.Enabled)
	r.ChunkPos(&x.ChunkRadius)
	r.Bool(&x.DoubleChests)
}
