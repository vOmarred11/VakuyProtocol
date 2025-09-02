package particle

import (
	"github.com/vOmarred11/VakuyProtocol/pos"
)

const (
	CauseEffect = iota
	CausePotion
)

// Particle is the default particle
type Particle struct {
	// Cause is the cause of particle release
	Cause uint8
	// Intensity is the release intensity
	Intensity int32
	// Position is the floating entity position
	Position pos.EntityPosition
}

func (p Particle) ParticleCause() uint8 {
	return p.Cause
}
func (p Particle) ParticleIntensity() int32 {
	return p.Intensity
}
func (p Particle) ParticlePosition() pos.EntityPosition {
	return p.Position
}
