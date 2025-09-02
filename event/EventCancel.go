package event

// Cancel cancel certain events
type Cancel struct {
	// Event is the chosen event
	Event Event
	// IsCancellable defines if the event can be cancelled
	IsCancellable bool
}

func (c Cancel) CancelEvent() Event {
	return c.Event
}
func (c Cancel) CancelIsCancellabile() bool {
	if c.IsCancellable == false {
		panic("you cannot cancel this event")
	}
	return c.IsCancellable
}
