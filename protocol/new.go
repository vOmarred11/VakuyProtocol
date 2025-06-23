package protocol

// New is every new value incoming
type New struct {
	// Type defines the type of the incoming value
	Type any
}

func (n New) NewType() any {
	return n.Type
}
func (n New) New(a any) any {
	return a
}
