package timeout

import "time"

// Time defines any duration sent by server
type Time struct {
	// Stamp is the time when you join the server
	Stamp time.Time
	// Sleep is the additional duration time when you join the server
	Sleep time.Duration
}

func (t Time) TimeStamp() time.Time {
	return t.Stamp
}
func (t Time) TimeSleep() time.Duration {
	return t.Sleep
}
