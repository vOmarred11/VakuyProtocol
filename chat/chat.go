package chat

import "fmt"

// Chat ...
type Chat struct {
	// Display is the default chat display
	Display string
	// FromServer defines if the message is sent by the server
	FromServer bool
}

// Entity is when the entity send a message
type Entity struct {
	// Name is the name of the entity
	// if the message is sent from the server it will be [SERVER]
	Name string
	// Message is the message from the entity
	Message string
	// MessageId is the id of the message
	MessageId int32
}

func (c Chat) ChatDisplay() string {
	e := Entity{}
	c.Display = fmt.Sprintf("<%s> %s", e.Name, e.Message)
	return c.Display
}
