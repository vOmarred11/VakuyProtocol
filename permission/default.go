package permission

// PermissionDefault returns the default permission level of each player
type PermissionDefault struct {
	// Level is the permission level
	Level uint8
}

func (p PermissionDefault) PermissionDefaultLevel() uint8 {
	return p.Level
}
