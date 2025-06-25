package area

type Position struct {
	AreaPosStart Area
	AreaPosEnd   Area
	error
}

func (pos Position) PositionAreaPosStart() Area {
	return pos.AreaPosStart
}
func (pos Position) PositionAreaPosEnd() Area {
	return pos.AreaPosEnd
}
func (pos Position) Error() error {
	return pos.error
}
