package pk

// Walk ... TODO figure out whats walk
type Walk struct {
	// Speed is the walking speed
	Speed float32
}

func (s Walk) WalkSpeed() float32 {
	return s.Speed
}
