package pos

import (
	"github.com/df-mc/dragonfly/server/world"
)

// EntityPosition ...
type EntityPosition struct {
	Type     world.EntityType
	Position float32
}

func (e EntityPosition) EntityType() world.EntityType {
	return e.Type
}
