package time

const (
	Day      = 1000
	Night    = 13000
	Noon     = 6000
	Midnight = 18000
	Sunrise  = 23000
	Sunset   = 12000
)

// Time is the time on the world
type Time struct {
	// Amount defines the time amount
	Amount uint8
}

func (t Time) TimeAmount() uint8 {
	return t.Amount
}
