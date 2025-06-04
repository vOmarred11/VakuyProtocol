package form

// ModalForm is another type of form
type ModalForm struct {
	// Title is the title of the ModalForm
	Title string
	// Content is a string that will be shown right down the title
	Content string
	// Button1 is the first button that will be shown
	Button1 struct {
		// Text is the text of the button
		Text string
		// Id is the id of the button which is always 1
		Id int32
		// Interaction is when the player clicks the button
		Interaction bool
	}
	// Button2 is the second button that will be shown
	Button2 struct {
		// Text is the text of the button
		Text string
		// Id is the id of the button which is always 1
		Id int32
		// Interaction is when the player clicks the button
		Interaction bool
	}
}

func (f *ModalForm) ModalFormTitle() string {
	return f.Title
}
func (f *ModalForm) ModalFormContent() string {
	return f.Content
}
func (f *ModalForm) ModalFormButton1Text() string {
	return f.Button1.Text
}
func (f *ModalForm) ModalFormButton1Id() int32 {
	return f.Button1.Id
}
func (f ModalForm) ModalFormButton1Interaction() bool {
	return f.Button2.Interaction
}
func (f *ModalForm) ModalFormButton2Text() string {
	return f.Button2.Text
}
func (f *ModalForm) ModalFormButton2Id() int32 {
	return f.Button2.Id
}
func (f *ModalForm) ModalFormButton2Interaction() bool {
	return f.Button2.Interaction
}
