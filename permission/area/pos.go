package area

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type Position struct {
	AreaPosStart float32
	AreaPosEnd   float32
	error
}

func (pos Position) PositionAreaPosStart() float32 {
	x := protocol.BlockPos{}.X()
	pos.AreaPosStart = float32(x) + 1
	return pos.AreaPosStart
}
func (pos Position) PositionAreaPosEnd() float32 {
	x := protocol.BlockPos{}.X()
	pos.AreaPosStart = float32(x) - 1
	return pos.AreaPosEnd
}
func (pos Position) Error() error {
	return pos.error
}
