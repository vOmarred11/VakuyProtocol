package pos

import (
	"github.com/df-mc/dragonfly/server/world"
)

// Entity TODO no need to explain, putting Entity to make the program happy
// Entity ...
type Entity struct {
	Type     world.EntityType
	Position float32
}

func (e Entity) EntityType() world.EntityType {
	return e.Type
}
