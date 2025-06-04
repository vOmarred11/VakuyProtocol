package panic

const (
	Max = 10000
)

// OverWriter returns an error when with the chosen max gets exceed
type OverWriter struct {
	// Amount is the amount of bytes
	Amount uint32
	// Type is the type of the byte
	Type byte
	// Overwrite returns an error
	Overwrite error
}

func (o *OverWriter) OverWrite() error {
	if Max < o.Amount {
		return nil
	}
	err := Panic("error overwrite: ", Error{}.Error)
	if err != nil {
		return err
	}
	return o.Overwrite
}
