package pk

import (
	"github.com/go-gl/mathgl/mgl32"
)

// ModernHystoy is when a player moves diagonally-horizontally
type ModernHystoy struct {
	// Angle is the angle of his move
	Angle mgl32.Vec3
	// Decay is the difference between the X and the Z
	Decay float32
	// BlockState is all blocks that the player touches on his walk
	BlockStateAmount int32
}

func (m ModernHystoy) ModernHystoyAngle() mgl32.Vec3 {
	return m.Angle
}
func (m ModernHystoy) ModernHystoyDecay() float32 {
	return m.Decay
}
func (m ModernHystoy) ModernHystoyBlockStateAmount() int32 {
	return m.BlockStateAmount
}
