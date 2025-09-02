package panic

type UnknownByte struct {
	Type byte
}

func (u UnknownByte) Unknown() error {
	if u.Type == 0 {
		Panic("returning an invalid value: type byte should be non-zero")
	}
	return nil
}
