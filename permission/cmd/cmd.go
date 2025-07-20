package cmd

import "github.com/vOmarred11/VakuyProtocol/player"

// PermissionCommand defines the command permission for each permission level
type PermissionCommand struct {
	// Level is the permission level
	Level uint8
	// Commands is a slice of commands that will have a permission
	Commands []string
	// HasPermission defines if the player has permissions
	HasPermission bool
}

func (p PermissionCommand) PermissionCommandLevel() uint8 {
	return p.Level
}
func (p PermissionCommand) PermissionCommandCommands() []string {
	return p.Commands
}
func (p PermissionCommand) PermissionCommandHasPermission() bool {
	pp := player.Player{}
	if p.HasPermission == false {
		pp.Message("You have no permission to use this command")
	}
	return p.HasPermission
}
