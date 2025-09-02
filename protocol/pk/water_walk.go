package pk

import (
	"github.com/vOmarred11/VakuyProtocol/pos"
)

// WaterWalk is when a player walks on water
// used for anticheats
type WaterWalk struct {
	// Speed is the walking speed
	Speed float32
	// Surface is the water surface
	Surface pos.BlockPos
}

func (w WaterWalk) WaterWalkSpeed() float32 {
	return w.Speed
}
func (w WaterWalk) WaterWalkSurface() pos.BlockPos {
	return w.Surface
}
