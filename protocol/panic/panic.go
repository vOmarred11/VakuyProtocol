package panic

// Error returns trivial errors
type Error struct {
	// Message is the message of the error
	Message string
	// Error returns an error as type error
	Error error
}

func Panic(e ...any) error {
	err := Panic(e)
	if err != nil {
		panic(err)
	}
	return err
}
func (e Error) PanicMessage() string {
	return e.Message
}
func (e Error) PanicError() error {
	return e.Error
}
