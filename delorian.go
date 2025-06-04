package delorian

import (
	"VakuyProtocol/delpacket"
	"VakuyProtocol/worlds"
	"github.com/sandertv/gophertunnel/minecraft/resource"
)

func WriteVakuyPacket(pk delpacket.DelPacket) uint32 {
	if pk.ID() == 0 {
		panic("invalid packet id")
	}
	return pk.ID()
}
func ReadVakuyPacket(pk delpacket.Delorian) uint32 {
	if pk.StringError() != nil {
		panic(pk.StringError())
	}
	return pk.ID()
}
func ReadPackets() (error, map[uint32]delpacket.DelPacket) {
	return nil, map[uint32]delpacket.DelPacket{}
}

func ResourcesPacks() resource.Pack {
	return resource.Pack{}
}
func World() worlds.ServerResponse {
	return worlds.ServerResponse{}
}
