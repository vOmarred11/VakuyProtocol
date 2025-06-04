package event

import "VakuyProtocol/player"

// ChatEvent when a player is chatting
type ChatEvent struct {
	// Message is the message from the player
	Message string
	// Format is the global chat format
	Format string
}
type ChatData struct {
	Player player.Player
}

func (c ChatEvent) ChatEventMessage() string {
	return c.Message
}
func (c ChatEvent) ChatEventFormat() string {
	return c.Format
}
func (d ChatData) ChatEventPlayer() player.Player {
	return d.Player
}
