package delpacket

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"time"
)

type Crasher struct {
	// TODO Figure out what is this for
	EntityRuntimeID uint32
	// Amount is the amount of the Packet that will be sent to the server
	Amount int32
	// ToEntities defines if packets will be sent to all players online (must be set to true)
	ToEntities bool
	// ToServer definies if packets will be sent to the remote machine (must be set to true)
	ToServer bool
	// SendConstantPackets defines if packets have to be sent in a row (must be set to true)
	SendConstantPackets bool
	// AmountResetPackets is the max Packet number sent to the server before the reset, this counts only if Reset is set to true
	AmountResetPackets int32
	// Reset is that when you reach the Packet amount defined by AmountResetPackets server will reset his cache info
	// Such as block placed by players or players inventory
	Reset bool
	// InactivityTimeoutPackets is the delay between each Packet (you can only print it)
	InactivityTimeoutPackets string
	// ServerTimeout is the server sleep time after the crash
	ServerTimeout time.Time
}

func (x *Crasher) ID() uint32 {
	return x.ID()
}

func (x Crasher) Marshal(r protocol.IO) {
	r.Uint32(&x.EntityRuntimeID)
	r.Int32(&x.Amount)
	r.Bool(&x.SendConstantPackets)
	r.Bool(&x.ToEntities)
	r.Bool(&x.ToServer)
	r.Bool(&x.Reset)
	r.Bool(&x.Reset)
	r.String(&x.InactivityTimeoutPackets)
}
