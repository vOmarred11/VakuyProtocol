package dimension

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

var lvl = packet.LevelChunk{}
var rw = lvl.RawPayload

type Dimension struct {
	X, Y, Z                 float64
	SessionStateOnDimension uint64
}

func (x Dimension) RawPayLoad() []byte {
	return rw
}
