package player

type Position struct {
	X float64
	Y float64
	Z float64
}

func (p Position) PositionX() float64 {
	return p.X
}
func (p Position) PositionY() float64 {
	return p.Y
}
func (p Position) PositionZ() float64 {
	return p.Z
}
