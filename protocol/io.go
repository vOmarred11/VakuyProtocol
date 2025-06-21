package protocol

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"reflect"
)

// Proto is the current marshaller protocol
type Proto struct {
	// Amount is the amount of packet that got sent
	Amount int32
	// Error returns an error
	Error error
	// Bytes is a slice of all bytes
	Bytes []byte
	// Reader is the default ghopertunnel reader
	Reader protocol.Reader
}

func (p *Proto) ProtocolBytes() []byte {
	return p.Bytes
}
func (p *Proto) Read(b []byte) (n int32, err error) {
	if len(b) == 0 {
		panic("empty byte slice: nil pointer interface or invalid packet reading")
	}
	return 0, nil
}
func (p *Proto) Write(b []byte) (n int32, err error) {
	if len(b) == 0 {
		panic("empty byte slice: nil pointer interface or invalid packet writing")
	}
	return p.Read(b)
}
func (p *Proto) ProtocolError() error {
	return p.Error
}
func (p *Proto) ID() uint32 {
	return p.ID()
}
func (p *Proto) Version() string {
	return CurrentVersion
}
func (p *Proto) Protocol() int {
	return CurrentProtocol
}
func (p *Proto) WriteByte(b byte) error {
	if b == 0 {
		panic("empty byte slice writing: bytes should be non-zero")
	}
	err := p.Error
	if err != nil {
		panic(err)
	}
	p.Bytes = append(p.Bytes, b)
	return err
}
func (p *Proto) ReadByte(b byte) (byte, error) {
	if len(p.Bytes) == 0 {
		panic("empty byte slice reading: bytes should be non-zero")
	}
	err := p.Error
	if err != nil {
		panic(err)
	}
	v := p.Bytes[p.Amount]
	b = v
	return v, err
}
func (p *Proto) SendValue(v any) (any, error) {
	if v == "" {
		panic("unknown type string allowed only type: byte, packet.Packet")
	}
	if v == 0 {
		panic("unknown type int/uint allowed only type: byte, packet.Packet")
	}
	if v == 0.0 {
		panic("unknown type float32/64 allowed only type: byte, packet.Packet")
	}
	if v == nil {
		panic("nil send value: value should be non-zero")
	}
	err := p.Error
	if err != nil {
		panic(err)
	}
	v = reflect.ValueOf(v)
	return v, nil
}
func (p *Proto) ReceiveValue(v any) error {
	if v == nil {
		panic("nil send value: value should be non-zero")
	}
	err := p.Error
	if err != nil {
		panic(err)
	}
	v = reflect.ValueOf(v)
	return nil
}
