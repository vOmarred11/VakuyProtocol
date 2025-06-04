package delpacket

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Packet represents a Packet that may be sent over a Minecraft network connection. The Packet needs to hold
// a method to encode itself to binary and decode itself from binary.
type DelPacket interface {
	// ID returns the ID of the Packet. All of these identifiers of packets may be found in id.go.
	ID() uint32
	// Marshal encodes or decodes a Packet, depending on the protocol.IO
	// implementation passed. When passing a protocol.Writer, Marshal will
	// encode the Packet into its binary representation and write it to the
	// protocol.Writer. On the other hand, when passing a protocol.Reader,
	// Marshal will decode the bytes from the reader into the Packet.
	Marshal(io protocol.IO)
}
type Delorian interface {
	ID() uint32
	// Comparabile compares a value type Byte in uint32 or uint64
	Comparabile() int32
	// Error returns an error
	Error() error
	// StringError returns an error as a string
	StringError() map[error]string
}
