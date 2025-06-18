package event

import (
	"github.com/vOmarred11/VakuyProtocol/dimension"
	"github.com/vOmarred11/VakuyProtocol/player"
)

// QuitEvent is when a player leave the game
type QuitEvent struct {
	// Position is the last player position of leave
	Position dimension.Data
}
type QuitData struct {
	Player player.Player
}

func (j QuitEvent) QuitEventPosition() dimension.Data {
	return j.Position
}
func (j QuitData) QuitDataPlayer() player.Player {
	return j.Player
}
