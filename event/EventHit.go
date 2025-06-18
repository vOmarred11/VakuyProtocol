package event

import "github.com/vOmarred11/VakuyProtocol/player"

// HitEvent is when a player hits any entity
type HitEvent struct {
	// Distance is the distance where the player hit the entity
	Distance float32
}
type HitData struct {
	Player player.Player
}

func (h *HitEvent) HitEventDistance() float32 {
	return h.Distance
}
func (d *HitData) HitEventPlayer() player.Player {
	return d.Player
}
