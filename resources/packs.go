package delorian

import (
	"github.com/sandertv/gophertunnel/minecraft/resource"
)

type Resource struct{ string }

func Pack(r ResourcePack) (resource.Pack, bool) {
	return r.Name, false
}
