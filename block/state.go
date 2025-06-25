package block

// ActionStates is used for the fill command
type ActionStates struct {
	// Replace replaces everything
	Replace bool
	// Fill fills the empty space but it doesn't destroy
	Fill bool
	// Destroy destroys block with drop
	Destroy bool
}

func (s ActionStates) ActionStatesReplace() bool {
	return s.Replace
}
func (s ActionStates) ActionStatesFill() bool {
	return s.Fill
}
func (s ActionStates) ActionStatesDestroy() bool {
	return s.Destroy	
}
