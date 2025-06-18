package event

import "github.com/vOmarred11/VakuyProtocol/player"

const (
	ExhaustCauseWalk = iota
	ExhaustCauseSprint
	ExhaustCauseJumpSprint
)

// ExhaustEvent is when a player get lower saturation
type ExhaustEvent struct {
	// Amount is the saturation amount
	Amount int32
	// Cause is the cause which the player got that decline of saturation
	Cause uint8
}
type ExhaustData struct {
	Player player.Player
}

func (e *ExhaustEvent) ExhaustEventCause() uint8 {
	return e.Cause
}
func (e *ExhaustEvent) ExhaustEventAmount() int32 {
	return e.Amount
}
func (d *ExhaustData) ExhaustEventPlayer() player.Player {
	return d.Player
}
