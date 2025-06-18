package event

import (
	"crypto/md5"
	"github.com/pelletier/go-toml"
	"github.com/vOmarred11/VakuyProtocol/player"
)

// Event returns a value for all events
type Event struct {
	// Player is the player that will execute the event
	Player player.Player
	// Id returns the event id
	Id [16]byte
	// Error returns an error
	Error error
}

func (e Event) EventPlayer() player.Player {
	return e.Player
}
func (e Event) EventId() [16]byte {
	b := e.Id == [16]byte{}
	if b {
		return [16]byte{}
	}
	md, err := toml.Marshal(e.Id)
	if err != nil {
		panic(err)
	}
	if len(md) != 16 {
		return [16]byte{}
	}
	return md5.Sum(md)
}
func (e Event) EventError() error {
	return e.Error
}
