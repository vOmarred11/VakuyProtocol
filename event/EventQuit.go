package event

import (
	"github.com/vOmarred11/VakuyProtocol/dimension"
	"github.com/vOmarred11/VakuyProtocol/player"
)

type QuitData struct {
	Player player.Player
}

func (j QuitData) QuitDataPlayer() player.Player {
	return j.Player
}
