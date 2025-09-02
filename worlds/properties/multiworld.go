package properties

import "github.com/df-mc/dragonfly/server/world"

func World(worlds *world.World) string {
	return worlds.Name()
}
func MultiWorld() uint {
	return 0
}

// Vakuy defines if when you switch world proxy should change values
func Vakuy() bool {
	return false
}
