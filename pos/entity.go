package pos

import "github.com/vOmarred11/VakuyProtocol/entity"

// Entity TODO no need to explain, putting Entity to make the program happy
// Entity ...
type Entity struct {
	Type entity.Type
}

func (e Entity) EntityType() entity.Type {
	return e.Type
}
