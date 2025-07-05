package backup

import (
	"github.com/pelletier/go-toml"
	"github.com/vOmarred11/VakuyProtocol/block"
	"github.com/vOmarred11/VakuyProtocol/item"
	"github.com/vOmarred11/VakuyProtocol/minecraft/protocol/packet"
	"github.com/vOmarred11/VakuyProtocol/player"
	"os"
)

var List []string

// Backup is a way to bring back the server to the state when backup was made
// if it will be used on a player it will change only his data
type Backup struct {
	// Data is the data of the server
	Data []byte
	// ItemData are all items
	ItemData item.Item
	// BlockData are all blocks
	BlockData block.BlockFactory
	// DimensionData is the data world of the dimension
	DimensionData packet.DimensionData
	// PlayerData is the data of the player
	PlayerData player.Player
}

func (b Backup) BackupData() []byte {
	return b.Data
}
func (b Backup) BackupItemData() item.Item {
	return b.ItemData
}
func (b Backup) BackupBlockData() block.BlockFactory {
	return b.BlockData
}
func (b Backup) BackupDimensionData() packet.DimensionData {
	return b.DimensionData
}
func (b Backup) BackupPlayerData() player.Player {
	return b.PlayerData
}

// NewBackup creates a new backup
func (b Backup) NewBackup(name string, loc string) []byte {
	var Srv []byte
	for c := range List {
		if List[c] == name {
			panic("ERROR #trace SRC | cannot duplicate backups")
		}
		bl, err := toml.Marshal(b.BlockData)
		if err != nil {
			panic(err)
		}
		Srv = append(Srv, bl...)
		err = os.WriteFile(loc, bl, 0600)
		if err != nil {
			panic(err)
		}
		it, err := toml.Marshal(b.ItemData)
		if err != nil {
			panic(err)
		}
		Srv = append(Srv, it...)
		err = os.WriteFile(loc, it, 0600)
		if err != nil {
			panic(err)
		}
		dm, err := toml.Marshal(b.DimensionData)
		if err != nil {
			panic(err)
		}
		Srv = append(Srv, dm...)
		err = os.WriteFile(loc, dm, 0600)
		if err != nil {
			panic(err)
		}
		player, err := toml.Marshal(b.PlayerData)
		if err != nil {
			panic(err)
		}
		Srv = append(Srv, player...)
		err = os.WriteFile(loc, player, 0600)
		if err != nil {
			panic(err)
		}
		List = append(List, name)
	}
	return Srv
}

// Restore applies the backup
func (b Backup) Restore(name string) error {
	var (
		ch  error
		loc string
		v   interface {
			Close() error
			Data(interface{}) byte
		}
	)
	for c := range List {
		if List[c] == name {
			r := toml.Unmarshal(b.NewBackup(name, loc), v)
			ch = r
			return r
		}
	}
	return ch
}
