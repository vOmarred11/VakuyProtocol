package pos

// Objective is the position of any object
type Objective struct {
	// Type is the object type
	Type byte
	// Position is the position where the object is located
	Position float32
}

func (o Objective) ObjectiveType() byte {
	return o.Type
}
func (o Objective) ObjectivePosition() float32 {
	x := BlockPos{}.X()
	o.Position = float32(x) + 1
	return o.Position
}
