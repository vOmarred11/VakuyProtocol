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
	a = a.(New)
	proto := Proto{}
	x := proto.Int()
	if a == x {
		panic("failed to create a new type: value shouldn't be type int")
	}
	if a == nil {
		panic("failed to create a new type: value shouldn't be type error")
	}
	return a
}
