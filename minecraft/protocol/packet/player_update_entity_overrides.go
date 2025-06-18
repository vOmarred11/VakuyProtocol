package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PlayerUpdateEntityOverridesTypeClearAll = iota
	PlayerUpdateEntityOverridesTypeRemove
	PlayerUpdateEntityOverridesTypeInt
	PlayerUpdateEntityOverridesTypeFloat
)

// PlayerUpdateEntityOverrides is sent by the server to modify an entity's properties individually.
type PlayerUpdateEntityOverrides struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// PropertyIndex is the index of the property to modify. The index is unique for each property of an entity.
	PropertyIndex uint32
	// Type is the type of action to perform with the property. It is one of the constants above.
	Type byte
	// IntValue is the new integer value of the property. It is only used when Type is set to
	// PlayerUpdateEntityOverridesTypeInt.
	IntValue int32
	// FloatValue is the new float value of the property. It is only used when Type is set to
	// PlayerUpdateEntityOverridesTypeFloat.
	FloatValue float32
}

// ID ...
func (*PlayerUpdateEntityOverrides) ID() uint32 {
	return IDPlayerUpdateEntityOverrides
}

func (pk *PlayerUpdateEntityOverrides) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varuint32(&pk.PropertyIndex)
	io.Uint8(&pk.Type)
	if pk.Type == PlayerUpdateEntityOverridesTypeInt {
		io.Int32(&pk.IntValue)
	} else if pk.Type == PlayerUpdateEntityOverridesTypeFloat {
		io.Float32(&pk.FloatValue)
	}
}
