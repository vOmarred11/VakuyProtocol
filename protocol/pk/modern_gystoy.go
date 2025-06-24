package pk

import (
	"github.com/go-gl/mathgl/mgl32"
)

// ModernGystoy is when a player moves diagonally-vertically
type ModernGystoy struct {
	// Angle is the angle of his move
	Angle mgl32.Vec3
	// Decay is the difference between the X and the Z
	Decay float32
	// BlockState is all blocks that the player touches on his walk
	BlockStateAmount int32
}

func (m ModernGystoy) ModernGystoyAngle() mgl32.Vec3 {
	return m.Angle
}
func (m ModernGystoy) ModernGystoyDecay() float32 {
	return m.Decay
}
func (m ModernGystoy) ModernGystoyBlockStateAmount() int32 {
	return m.BlockStateAmount
}
