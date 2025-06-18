package ncp

type Position struct {
	// SpawnPosition is the position where Ncp spawned
	SpawnPosition float32
	// Position is where the Ncp is located at this movement
	Position []float32
	// FinalPosition is the last Ncp position
	FinalPosition float32
}

func (p Position) NcpSpawnPosition() float32 {
	return p.SpawnPosition
}
func (p Position) NcpPosition() []float32 {
	if len(p.Position) == 3 {
		panic("too many arguments, shouldn't reach 3 (X-Y-Z)")
	}
	return p.Position
}
func (p Position) NcpFinalPosition() float32 {
	return p.FinalPosition
}
