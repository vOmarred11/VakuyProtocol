package text

import "github.com/vOmarred11/VakuyProtocol/pos"

type ParticleText struct {
	Text string
	// Intensity is the release intensity
	Intensity int32
	// Position is the floating entity position
	Position pos.Objective
}

func (p ParticleText) ParticleText() string {
	return p.Text
}
func (p ParticleText) ParticleTextIntensity() int32 {
	return p.Intensity
}
func (p ParticleText) ParticleTextPosition() pos.Objective {
	return p.Position
}
