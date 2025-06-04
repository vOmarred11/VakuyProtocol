package form

// SimpleForm is a basic form
type SimpleForm struct {
	// Title is the title of the SimpleForm
	Title string
	// Content is a string that will be shown right down the title
	Content string
	// Button is the button which will be sent
	Button struct {
		// Text is the text of the button
		Text string
		// Id is the id of the button
		// DISCLAMER: if the first id is 1 and second is 2, it will show the first one so 1
		Id int32
		// Interaction is when the player clicks the button
		Interaction bool
	}
}

func (x SimpleForm) SimpleFormTitle() string {
	return x.Title
}
func (x SimpleForm) SimpleFormContent() string {
	return x.Content
}
func (x SimpleForm) SimpleFormButtonText() string {
	return x.Button.Text
}
func (x SimpleForm) SimpleFormButtonId() int32 {
	return x.Button.Id
}
func (x SimpleForm) SimpleFormButtonInteraction() bool {
	return x.Button.Interaction
}
