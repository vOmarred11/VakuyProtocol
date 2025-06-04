package ncp

// Default is a struct of default values when Ncp gets created
type Default struct {
	// Hearts is the heart amount
	// This only works when Damage is active
	Hearts int32
	// Position is the position of the Ncp on spawn
	Position float32
}

func (d Default) DefaultHearts() int32 {
	return d.Hearts
}
func (d Default) DefaultPosition() float32 {
	return d.Position
}
