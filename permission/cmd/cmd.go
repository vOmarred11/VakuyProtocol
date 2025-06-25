package cmd

// PermissionCommand defines the command permission for each permission level
type PermissionCommand struct {
	// Level is the permission level
	Level uint8
	// Commands is a slice of commands that will have a permission
	Commands []string
}

func (p PermissionCommand) PermissionCommandLevel() uint8 {
	return p.Level
}
func (p PermissionCommand) PermissionCommandCommands() []string {
	return p.Commands
}
