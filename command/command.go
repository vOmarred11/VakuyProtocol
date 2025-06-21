package command

// Command is a message that you can execute in chat with the / to execute actions
type Command struct {
	// Name is the name of the command
	Name string
	// Description is the description that will appear
	Description string
	// PermissionLevel defines who can execute that command
	PermissionLevel int32
	// Aliases is a slice of all connections between the command
	Aliases []string
	// Args is a slice of all command arguments
	Args []Arguments
}

func (x Command) CommandName() string {
	return x.Name
}
func (x Command) CommandDescription() string {
	return x.Description
}
func (x Command) CommandAliases() []string {
	return x.Aliases
}
