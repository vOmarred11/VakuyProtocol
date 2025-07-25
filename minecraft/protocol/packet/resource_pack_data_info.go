package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// noinspection SpellCheckingInspection
const (
	ResourcePackTypeAddon = iota + 1
	ResourcePackTypeCached
	ResourcePackTypeCopyProtected
	ResourcePackTypeBehaviour
	ResourcePackTypePersonaPiece
	ResourcePackTypeResources
	ResourcePackTypeSkins
	ResourcePackTypeWorldTemplate
)

// ResourcePackDataInfo is sent by the server to the client to inform the client about the data contained in
// one of the resource pack that are about to be sent.
type ResourcePackDataInfo struct {
	// UUID is the unique ID of the resource pack that the info concerns.
	UUID string
	// DataChunkSize is the maximum size in bytes of the chunks in which the total size of the resource pack
	// to be sent will be divided. A size of 1MB (1024*1024) means that a resource pack of 15.5MB will be
	// split into 16 data chunks.
	DataChunkSize uint32
	// ChunkCount is the total amount of data chunks that the sent resource pack will exist out of. It is the
	// total size of the resource pack divided by the DataChunkSize field.
	// The client doesn't actually seem to use this field. Rather, it divides the size by the chunk size to
	// calculate it itself.
	ChunkCount uint32
	// Size is the total size in bytes that the resource pack occupies. This is the size of the compressed
	// archive (zip) of the resource pack.
	Size uint64
	// Hash is a SHA256 hash of the content of the resource pack.
	Hash []byte
	// Premium specifies if the resource pack was a premium resource pack, meaning it was bought from the
	// Minecraft store.
	Premium bool
	// PackType is the type of the resource pack. It is one of the resource pack types that may be found in
	// the constants above.
	PackType byte
}

// ID ...
func (*ResourcePackDataInfo) ID() uint32 {
	return IDResourcePackDataInfo
}

func (pk *ResourcePackDataInfo) Marshal(io protocol.IO) {
	io.String(&pk.UUID)
	io.Uint32(&pk.DataChunkSize)
	io.Uint32(&pk.ChunkCount)
	io.Uint64(&pk.Size)
	io.ByteSlice(&pk.Hash)
	io.Bool(&pk.Premium)
	io.Uint8(&pk.PackType)
}
