package delorian

import (
	"github.com/sandertv/gophertunnel/minecraft/resource"
)

type ResourcePack struct {
	string
}

func DownloadCache() string {
	return "delorian/packs"
}
func DownloadPacks() string {
	return "delorian/downloads"
}
func Pack(r Resources) (resource.Pack, bool) {
	return r.Name, false
}
func Resource() ResourcePack { return Resource() }
