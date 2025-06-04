package command

const (
	Arg = iota
	ArgCommandTypeString
	ArgCommandTypeInt
	ArgCommandTypeFloat
	ArgCommandTypeBool
	ArgCommandTypeCustomList
	ArgCommandTypeBlockStates
	ArgCommandTypeRawText
	ArgCommandTypeValid
)

// Arguments are command arguments
type Arguments struct {
	// Value is the string value that will be shown
	Value string
	// Arg is the arg type
	Arg uint
	// Optional basically changes nothing in terms of game, but it will show () instead of []
	Optional bool
}

func (x Arguments) CommandValue() string {
	return x.Value
}
func (x Arguments) CommandArg() uint {
	return x.Arg
}
func (x Arguments) CommandOptional() bool {
	return x.Optional
}
