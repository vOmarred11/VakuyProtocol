package delread

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Any reads any incoming packet
type Any struct {
	Read protocol.Reader
}

func (x *Any) ID() uint32 {
	return x.ID()
}

func (x *Any) Marshal() uint32 {
	return x.Marshal()
}
func (x *Any) Error() error {
	return x.Error()
}
func (x *Any) Comparabile() int32 {
	return x.Comparabile()
}
func (x *Any) StringError() map[error]string {
	return x.StringError()
}
