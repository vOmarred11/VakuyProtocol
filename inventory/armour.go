package inventory

const (
	ArmourHelmet = iota
	ArmourChestPlate
	ArmourLeggings
	ArmourBoots
)

// Armour ...
// TODO no need to explain, i think
type Armour struct {
	Helmet     InventoryData
	ChestPlate InventoryData
	Leggings   InventoryData
	Boots      InventoryData
}

func (a Armour) ArmourHelmet() InventoryData {
	return a.Helmet
}
func (a Armour) ArmourChestPlate() InventoryData {
	return a.ChestPlate
}
func (a Armour) ArmourLeggings() InventoryData {
	return a.Leggings
}
func (a Armour) ArmourBoots() InventoryData {
	return a.Boots
}
