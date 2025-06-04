package command

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Args        []Arguments
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
