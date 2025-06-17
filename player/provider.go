package player

import (
	"errors"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/google/uuid"
	"io"
)

// Provider represents a value that may provide data to a Player value. It usually does the reading and
// writing of the player data so that the Player may use it.
type Provider interface {
	// Save is called when the player leaves the server. The Data of the player is passed.
	Save(uuid uuid.UUID, data Config, w *world.World) error
	// Load is called when the player joins and passes the UUID of the player.
	// It expects to the player data, and an error that is nil if the player data could be found. If non-nil, the player
	// will use default values, and you can use an empty Data struct.
	Load(uuid uuid.UUID, world func(world.Dimension) *world.World) (Config, *world.World, error)
	// Closer is used on server close when the server calls Provider.Close() and is used to safely close the Provider.
	io.Closer
}

// Compile time check to make sure NopProvider implements Provider.
var _ Provider = (*NopProvider)(nil)

// NopProvider is a player data provider that won't store any data and instead always return default values
type NopProvider struct{}

func (NopProvider) Save(uuid.UUID, Config, *world.World) error { return nil }
func (NopProvider) Load(uuid.UUID, func(world.Dimension) *world.World) (Config, *world.World, error) {
	return Config{}, nil, errors.New("")
}
func (NopProvider) Close() error { return nil }
