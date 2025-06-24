package event

type Cancel struct {
	Event          Event
	IsCancellabile bool
}

func (c Cancel) CancelEvent() Event {
	return c.Event
}
func (c Cancel) CancelIsCancellabile() bool {
	if c.IsCancellabile == false {
		panic("you cannot cancel this event")
	}
	return c.IsCancellabile
}
