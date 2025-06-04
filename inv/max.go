package inv

const (
	InventoryMaxItem  = 64
	InventoryMaxBlock = 64
)

type InventoryMax struct {
	// Quantity is the max quantity for each block
	Quantity int32
	// ChangeForBlock defines if the max stack should change for each block
	// For example max dirt should be 70 and max stone should be 64
	ChangeForBlock bool
}

func (i InventoryMax) InventoryMaxQuantity() int32 {
	return i.Quantity
}
func (i InventoryMax) InventoryMaxChangeForBlock() bool {
	return i.ChangeForBlock
}
