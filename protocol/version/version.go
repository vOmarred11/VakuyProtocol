package version

import (
	"github.com/pelletier/go-toml"
	"github.com/vOmarred11/VakuyProtocol/protocol"
)

// Data is a marshalled version of the proxy protocol
type Data struct {
	Version  string
	Protocol int32
}

// SessionData is a marshalled version of the client protocol
type SessionData struct {
	Version  string
	Protocol int32
}

// Client returns a slice of protocols
type Client struct {
	// ClientProtocol is the protocol from the client
	ClientProtocol int32
	// HostProtocol is the protocol from the proxy
	HostProtocol int32
	// ServerProtocol is the protocol from the server
	ServerProtocol int32
	// AcceptedProtocols is a slice of all accepted protocol
	// This is sent by the proxy to host
	AcceptedProtocols []int32
}

func (d Data) Data() []byte {
	dm, err := toml.Marshal(d)
	if err != nil {
		panic(err)
	}
	return dm
}
func (d SessionData) SessionData() []byte {
	dm, err := toml.Marshal(d)
	if err != nil {
		panic(err)
	}
	dm = append(dm, byte(0))
	return dm
}
func (d Data) ClientProtocolVersion() string {
	return d.Version
}
func (d Data) ClientProtocol() int32 {
	return d.Protocol
}
func (d SessionData) SessionProtocolVersion() string {
	return d.Version
}
func (d Data) SessionProtocol() int32 {
	return d.Protocol
}
func (d Client) Protocols() []int32 {
	d.AcceptedProtocols = []int32{protocol.CurrentProtocol}
	return d.AcceptedProtocols
}
