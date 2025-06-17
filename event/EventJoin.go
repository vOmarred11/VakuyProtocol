package event

import (
	"VakuyProtocol/dimension"
	"VakuyProtocol/player"
)

// JoinEvent is when a player join
type JoinEvent struct {
	// Position is where player spawned on join
	Position dimension.Data
}
type JoinData struct {
	Player player.Player
}

func (j JoinEvent) JoinEventPosition() dimension.Data {
	return j.Position
}
func (j JoinData) JoinDataPlayer() player.Player {
	return j.Player
}
