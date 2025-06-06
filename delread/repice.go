package delread

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// DelPacket represents a Packet that may be sent over a Minecraft network connection. The Packet needs to hold
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
