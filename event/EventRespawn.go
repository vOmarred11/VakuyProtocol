package event

import (
	"VakuyProtocol/player"
)

// RespawnEvent is when a player respawns
type RespawnEvent struct {
	// Position is the respawn position point
	Position player.Position
}
type RespawnData struct {
	Player player.Player
}

func (r RespawnEvent) RespawnEventPosition() player.Position {
	return r.Position
}
func (d RespawnData) RespawnDataPlayer() player.Player {
	return d.Player
}
