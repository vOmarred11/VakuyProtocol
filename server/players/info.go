package players

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Info returns target information
type Info struct {
	// Name is the target name
	Name string
	// Version is the target version
	Version string
	// Protocol is the target protocol
	Protocol int32
	// Skin returns the target skin
	Skin protocol.Skin
	// DeviceModel is the device which target is on
	DeviceModel protocol.DeviceOS
}
