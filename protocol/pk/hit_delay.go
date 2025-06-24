package pk

// HitDelay is the default minecraft hit delay
type HitDelay struct {
	// Amount is the delay amount
	Amount float64
}

func (h HitDelay) HitDelayAmount() float64 {
	return h.Amount
}
