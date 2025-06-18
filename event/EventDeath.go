package event

import "github.com/vOmarred11/VakuyProtocol/player"

const (
	DeathCauseHit = iota
	DeathCauseFall
	DeathCauseFire
	DeathCauseSuffucation
	DeathCauseDrown
	DeathCauseEntity
)

// DeathEvent is when a player dies
type DeathEvent struct {
	// Cause is the death cause
	Cause uint8
	// Position is the position where player dies
	Position player.Position
}
type DeathData struct {
	Player player.Player
}

func (d DeathEvent) DeathEventCause() uint8 {
	return d.Cause
}
func (d DeathEvent) DeathEventPosition() player.Position {
	return d.Position
}
func (d DeathData) DeathEventPlayer() player.Player {
	return d.Player
}
