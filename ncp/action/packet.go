package action

const (
	NcpAction = iota
	NcpActionStart
	NcpActionStop

	NcpActionStartWalking
	NcpActionStopWalking
	NcpActionStartJumping
	NcpActionStopJumping
	NcpActionStartMoveHead
	NcpActionStopMoveHead
	NcpActionStartSwim
	NcpActionStopSwim
	NcpActionStartJumpingSwim
	NcpActionStopJumpingSwim
	NcpActionStartSprint
	NcpActionStopSprint
	NcpActionStartFlying
	NcpActionStopFlying
	NcpActionStartFalling
	NcpActionStopFalling
	NcpActionStartWalkSpeed
	NcpActionStopWalkSpeed
	NcpActionStartAttacking
	NcpActionStopAttacking
	NcpActionStartBreakingBlocks
	NcpActionStopBreakingBlocks
	NcpActionStartPlacingBlocks
	NcpActionStopPlacingBlocks
	NcpActionDropItems
	NcpActionDropBlocks
	NcpActionStartThrowingItem
	NcpActionStopThrowingItem
	NcpActionStartBowing
	NcpActionStopBowing
	NcpActionStartBurning
	NcpActionStopBurning
	NcpActionOpenDoor
	NcpActionCloseDoor
)

func ActionId(id int) int {
	return id
}
