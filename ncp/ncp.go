package ncp

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	Skin = iota
	SkinTypeSteve
	SkinTypeDefault
)

type Ncp struct {
	// Name is the of the Ncp
	Name string
	// Skin is the skin of the Ncp
	Skin protocol.Skin
	// Interactable is a struct of all action that the player can do on the Ncp
	Interactable map[bool]struct {
		MouseRightClick bool
		MouseLeftClick  bool
	}
}

func (n Ncp) NcpName() string {
	return n.Name
}
func (n Ncp) NcpSkin() protocol.Skin {
	return n.Skin
}
func (n Ncp) NcpInteractableMouseRightClick() bool {
	return n.Interactable[].MouseRightClick
}
func (n Ncp) NcpInteractableMouseLeftClick() bool {
	return n.Interactable[].MouseLeftClick
}
