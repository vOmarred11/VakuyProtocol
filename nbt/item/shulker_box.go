package item

type ShulkerBox struct {
	Data []byte
}

func (e ShulkerBox) ShulkerBoxData() []byte {
	return e.Data
}
