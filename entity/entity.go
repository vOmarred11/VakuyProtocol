package entity

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/vOmarred11/VakuyProtocol/pos"
)

// Type is the entity type
type Type struct {
	// Behaviour returns the entity behaviour
	Behaviour byte
	// TODO do i need to explain
	Entity world.Entity
	// Position is the position of the entity
	Position *pos.Entity
}

func (t Type) TypeBehaviour() byte {
	return t.Behaviour
}
func (t Type) TypeEntity() world.Entity {
	return t.Entity
}
