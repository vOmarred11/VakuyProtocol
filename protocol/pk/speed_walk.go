package pk

// SpeedWalk is when a player is faster than the default minecraft speed
type SpeedWalk struct {
	// Speed is the player speed
	Speed float32
}

func (s SpeedWalk) SpeedWalkSpeed() float32 {
	return s.Speed
}
