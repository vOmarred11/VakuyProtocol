package byte

const (
	Message = 0x104281
)

// WriterByte in this case just sends a message, but you can do whatever you want just need codes
type WriterByte struct {
	Message uint8
}

func (s WriterByte) MessageByte() uint8 {
	return s.Message
}
