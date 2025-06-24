package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ResourcePackStack is sent by the server to send the order in which resource pack and behaviour pack
// should be applied (and downloaded) by the client.
type ResourcePackStack struct {
	// TexturePackRequired specifies if the client must accept the texture pack the server has in order to
	// join the server. If set to true, the client gets the option to either download the resource pack and
	// join, or quit entirely. Behaviour pack never have to be downloaded.
	TexturePackRequired bool
	// BehaviourPack is a list of behaviour pack that the client needs to download before joining the server.
	// All of these behaviour pack will be applied together, and the order does not necessarily matter.
	BehaviourPacks []protocol.StackResourcePack
	// TexturePacks is a list of texture pack that the client needs to download before joining the server.
	// The order of these texture pack specifies the order that they are applied in on the client side. The
	// first in the list will be applied first.
	TexturePacks []protocol.StackResourcePack
	// BaseGameVersion is the vanilla version that the client should set its resource pack stack to.
	BaseGameVersion string
	// Experiments holds a list of experiments that are either enabled or disabled in the world that the
	// player spawns in.
	// It is not clear why experiments are sent both here and in the StartGame packet.
	Experiments []protocol.ExperimentData
	// ExperimentsPreviouslyToggled specifies if any experiments were previously toggled in this world. It is
	// probably used for some kind of metrics.
	ExperimentsPreviouslyToggled bool
	// IncludeEditorPacks specifies if vanilla editor pack should be included in the resource pack stack when
	// connecting to an editor world.
	IncludeEditorPacks bool
}

// ID ...
func (*ResourcePackStack) ID() uint32 {
	return IDResourcePackStack
}

func (pk *ResourcePackStack) Marshal(io protocol.IO) {
	io.Bool(&pk.TexturePackRequired)
	protocol.Slice(io, &pk.BehaviourPacks)
	protocol.Slice(io, &pk.TexturePacks)
	io.String(&pk.BaseGameVersion)
	protocol.SliceUint32Length(io, &pk.Experiments)
	io.Bool(&pk.ExperimentsPreviouslyToggled)
	io.Bool(&pk.IncludeEditorPacks)
}
