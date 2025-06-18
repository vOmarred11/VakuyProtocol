package floating

import (
	"github.com/vOmarred11/VakuyProtocol/pos"
)

// Floating represent the floating entity
type Floating struct {
	// Type is the floating entity type
	Type byte
	// Intensity represent the intensity of the floating entity
	Intensity int32
	// Position is where the floating entity is located
	Position pos.Objective
}

func (f Floating) FloatingType() byte {
	return f.Type
}
func (f Floating) FloatingIntensity() int32 {
	return f.Intensity
}
func (f Floating) FloatingPosition() pos.Objective {
	return f.Position
}
