package hit

import "github.com/vOmarred11/VakuyProtocol/pos"

// ParticleCriticalHit represent particles when a player crit you
type ParticleCriticalHit struct {
	// Count is the number count of floating entities
	// Since this packet cannot be or rarely gets sent we are not calculating the intensity but the number
	Count int32
	// Position the position of each floating entity
	Position pos.Entity
}

func (p ParticleCriticalHit) ParticleCriticalHitCount() int32 {
	return p.Count
}
func (p ParticleCriticalHit) ParticleCriticalHitPosition() pos.Entity {
	return p.Position
}
