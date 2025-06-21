package timeout

import "time"

// Timeout defines any timeout sent by client
type Timeout struct {
	// Amount is the sleep time when you're joining the server
	Amount time.Duration
}

func (t Timeout) TimeoutAmount() time.Duration {
	return t.Amount
}
