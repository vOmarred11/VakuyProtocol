package cooldown

const (
	TypeEvent = iota
	TypeItemUse
)

// Cooldown is the time when you can't use that action
type Cooldown struct {
	// Amount is the seconds amount
	Amount int32
	// Type is the type of the usage
	Type uint8
}

func (c Cooldown) CooldownAmount() int32 {
	return c.Amount
}
func (c Cooldown) CooldownType() uint8 {
	return c.Type
}
