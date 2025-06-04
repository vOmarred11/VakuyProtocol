package dimension

// Data defines world's manage struct
type Data struct {
	// ID returns the dimension id
	ID uint32
	// Name is the name of the dimension
	Name string
	// RawPayLoad is a dimension chunk data in byte
	RawPayLoad []byte
}

func (d Data) DimensionID() uint32 {
	return d.ID
}
func (d Data) DimensionName() string {
	return d.Name
}
func (d Data) RawPayload() []byte {
	return d.RawPayLoad
}
