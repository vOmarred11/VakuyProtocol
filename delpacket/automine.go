package delpacket

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	DiamondOre = iota
	GoldenOre
	IronOre
	CoalOre
	EmeraldOre
	RedstoneOre
	LapisOre
)
const (
	KeyTypeMiddleClick = iota
)

// AutoMine automatically teleports you to chosen blocks
type AutoMine struct {
	// TODO: figure out what this is for
	EntityRuntimeID uint64
	// If AvoidLava is set to true, it will skip ores next to lava
	AvoidLava bool
	// If AvoidNonClusters is set to true, it will skip non-clusters, so 1/2 diamond ore block veins
	AvoidNonClusters bool
	// If Enabled is set to true, every time you click the chosen key, it will teleport you to the chosen ore
	Enabled bool
	// If Pickaxe is set to true, you need a pickaxe to teleport to the chosen ore
	Pickaxe bool
	// Range is the range within which you can teleport to blocks
	Range float32
	// Ore is the type of ore you will be teleported to
	Ore uint64
	// Mode is the teleportation mode
	Mode byte
	// Read is the reading module
	Read interface{}
}

func (x *AutoMine) ID() uint32 {
	return x.ID()
}

func (x *AutoMine) Marshal(r protocol.IO) {
	r.Bool(&x.AvoidLava)
	r.Bool(&x.AvoidNonClusters)
	r.Bool(&x.Enabled)
	r.Bool(&x.Pickaxe)
	r.Float32(&x.Range)
	r.Uint64(&x.Ore)
	r.Uint8(&x.Mode)
	r.Uint64(&x.EntityRuntimeID)
}
