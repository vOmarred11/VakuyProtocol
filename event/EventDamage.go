package event

import "VakuyProtocol/player"

const (
	DamageCauseHit = iota
	DamageCauseFall
	DamageCauseFire
	DamageCauseSuffucation
	DamageCauseDrown
	DamageCauseEntity
)

// DamageEvent is when a player takes damage
type DamageEvent struct {
	// Amount is the damage amount
	Amount int32
	// Cause is the damage cause
	Cause uint8
}
type DamageData struct {
	Player player.Player
}

func (d DamageEvent) DamageEventAmount() int32 {
	return d.Amount
}
func (d DamageEvent) DamageEventCause() uint8 {
	return d.Cause
}
func (d DamageData) DamageEventPlayer() player.Player {
	return d.Player
}
