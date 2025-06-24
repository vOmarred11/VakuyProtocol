package pk

// AirJump is when a player jumps in the air
type AirJump struct {
	// Height is the height where player is located before and after the jump
	Height []float32
}

func (a AirJump) AirJumpHeight() ([]float32, []float32) {
	initial := append(a.Height, a.Height[0])
	final := append(a.Height, a.Height[1])
	return initial, final
}
