package pk

// Jump ... TODO figure out whats jump
type Jump struct {
	// Height is the height of the jump
	Height float32
}

func (j Jump) JumpHeight() float32 {
	return j.Height
}
