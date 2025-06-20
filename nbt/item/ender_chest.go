package item

type EnderChest struct {
	Data []byte
}

func (e EnderChest) EnderChestData() []byte {
	return e.Data
}
