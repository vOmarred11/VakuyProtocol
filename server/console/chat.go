package console

import (
	"github.com/vOmarred11/VakuyProtocol/chat"
)

// Chat is the output on the console
type Chat struct {
	// ConsoleDisplay is the display that will be shown on console
	ConsoleDisplay string
	// CensureBadWords defines if also in the console chat bad words should be censured
	CensureBadWords bool
}

func (c *Chat) ChatConsoleDisplay() string {
	ch := chat.Chat{}
	ch.Display = c.ConsoleDisplay
	return c.ConsoleDisplay
}
func (c *Chat) ChatCensureBadWords() bool {
	r := chat.Restriction{}
	if c.CensureBadWords == true {
		r.Enabled = true
	}
	return c.CensureBadWords
}
