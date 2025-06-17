package chat

// BadWords is an array of all bad words
var BadWords []string

// Restriction restricts all bad words censuring them
type Restriction struct {
	// ErrorMessage is the message that will be shown when someone will type a bad word
	ErrorMessage string
	// Message is the message containing bad words
	Message string
	// Level is the restriction level
	Level byte
}

func (r Restriction) RestrictionErrorMessage() string {
	return r.ErrorMessage
}
func (r Restriction) RestrictionMessage() string {
	return r.Message
}
func (r Restriction) RestrictionLevel() byte {
	return r.Level
}
func RestrictionBadWords() []string {
	return BadWords
}
