package time

// Spawn defines what entities will spawn in the selected time
// if you are using this on the proxy it won't work
type Spawn struct {
	// Time is the selected item
	Time uint8
	// Mob defines if mob will spawn on the selected time
	Mob bool
	// Animals defines if animals will spawn on the selected time
	Animals bool
}

func (s Spawn) SpawnTime() uint8 {
	return s.Time
}
func (s Spawn) SpawnMob() bool {
	return s.Mob
}
func (s Spawn) SpawnAnimals() bool {
	return s.Animals
}
