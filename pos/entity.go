package pos

import (
	"github.com/df-mc/dragonfly/server/world"
)

// EntityPosition ...
type EntityPosition struct {
	// Type is the world where player is located at
	World world.EntityType
	// Position is the exact position
	Position float32
}

func (e EntityPosition) EntityType() world.EntityType {
	return e.World
}
