package ohc

const (
	ActionGiveHealth = iota
	ActionGiveSaturation
	ActionGiveBlock
	ActionGiveItem
	ActionMove
	ActionChangeInventory
	ActionDisconnect
	ActionKill
	ActionMessage
	ActionTeleport
)

// Action is when the server sends an action to the player
type Action struct {
	// Type is the action
	Type uint8
}

func (a Action) ActionType() uint8 {
	return a.Type
}
