package players

// Hit defines when a player hits you
type Hit struct {
	// Distance is the block distance
	Distance float32
	// OnHitbox defines if the hit is on the player skin or outside
	OnHitbox bool
}

func (h *Hit) HitDistance() float32 {
	return h.Distance
}
func (h *Hit) HitOnHitbox() bool {
	return h.OnHitbox
}
