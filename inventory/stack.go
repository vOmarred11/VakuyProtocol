package inventory

// Stack is the quantity of each block
type Stack struct {
	// Block returns the which item get changed
	Block uint8
	// Max returns the max of the block
	Max int32
}

func (s Stack) StackBlock() uint8 {
	return s.Block
}
func (s Stack) StackMax() int32 {
	return s.Max
}
