package command

// Output is the action when the command gets executed
type Output struct {
	// TODO figure out what is Action for
	Action any
	// Player defines if the packet should be sent to the player or to the server
	// An example can be broadcast messages which is sent to the whole server
	Player bool
}

func (x Output) CommandAction() any {
	return x.Action
}
func (x Output) CommandPlayer() bool {
	return x.Player
}
