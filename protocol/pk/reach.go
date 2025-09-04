package pk

const (
	ReachGamemodeSurvival  = 3
	ReachGamemodeAdventure = 3
	ReachGamemodeCreative  = 7

	ReachBlock = 4.5
)

// Reach is the minecraft hit reach
type Reach struct {
	// Default is the default minecraft hit reach
	Default float32
	// Gamemode defines each gamemode reach
	Gamemode uint8
}

// BlockReach is the minecraft block interaction reach
type BlockReach struct {
	// Default is the default value of the block reach which is 4.5
	Default float32
}

func (r Reach) ReachDefault() float32 {
	return r.Default
}
func (r Reach) ReachGamemode() uint8 {
	return r.Gamemode
}
func (r Reach) BlockReachDefault() float32 {
	return r.Default
}
