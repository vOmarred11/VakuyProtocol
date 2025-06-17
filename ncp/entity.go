package ncp

type Entity struct {
	Attack     bool
	Damage     bool
	UseItems   bool
	Walk       bool
	Jump       bool
	Sprint     bool
	DropItems  bool
	TakeItems  bool
	PlaceBlock bool
	BreakBlock bool
	Swim       bool
	OpenChest  bool
	OpenDoor   bool
	PickUpItem bool
	ThrowItems bool
	Burn       bool
}

func (e Entity) NcpAttack() bool {
	return e.Attack
}
func (e Entity) NcpDamage() bool {
	return e.Damage
}
func (e Entity) NcpUseItems() bool {
	return e.UseItems
}
func (e Entity) NcpWalk() bool {
	return e.Walk
}
func (e Entity) NcpJump() bool {
	return e.Jump
}
func (e Entity) NcpSprint() bool {
	return e.Sprint
}
func (e Entity) NcpDropItems() bool {
	return e.DropItems
}
func (e Entity) NcpTakeItems() bool {
	return e.TakeItems
}
func (e Entity) NcpPlaceBlock() bool {
	return e.PlaceBlock
}
func (e Entity) NcpBreakBlock() bool {
	return e.BreakBlock
}
func (e Entity) NcpSwim() bool {
	return e.Swim
}
func (e Entity) NcpOpenChest() bool {
	return e.OpenChest
}
func (e Entity) NcpOpenDoor() bool {
	return e.OpenDoor
}
func (e Entity) NcpThrowItems() bool {
	return e.ThrowItems
}
func (e Entity) NcpBurn() bool {
	return e.Burn
}
