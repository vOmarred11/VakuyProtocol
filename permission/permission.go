package permission

const (
	PermissionLevelMember = iota
	PermissionLevelVisitator
	PermissionLevelCustom
	PermissionLevelOperator
)

// PermissionLevel returns the level of the permission
type PermissionLevel struct {
	// Level is the permission level declared by const statement
	Level uint8
}

func (p PermissionLevel) PermissionLevel() uint8 {
	return p.Level
}
