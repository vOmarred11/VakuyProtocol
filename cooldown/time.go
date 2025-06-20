package cooldown

import "time"

// Time is the sleep time
type Time struct {
	Time time.Duration
}

func (t Time) TimeDuration() time.Duration {
	return t.Time
}
