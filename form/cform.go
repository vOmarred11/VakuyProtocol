package form

// CustomForm is a variable and more versatile form
// The button is always "Submit"
type CustomForm struct {
	// Title is the title of the CustomForm
	Title string
	// Content is a string that will be shown right down the Title
	Content string
	// TextBox is an area where you can type
	// By adding a text in you will automatically add also the TextBox and the text added on the string
	// will be transparent
	TextBox string
	// Slider is a slider where you can choose an element
	Slider struct {
		// Text of the Slider
		Text string
		// The max of the Slider
		Max int32
		// The min of the Slider
		Min int32
	}
	// Dropdown is an array of argument
	Dropdown struct {
		// Text is the text of the dropdown
		Text string
		// Arguments is an array of all arguments
		Arguments []string
	}
	// Toggle is a toggler where you can choose if its true or false
	Toggle struct {
		// Text is the text of the Toggle
		Text string
		// Toggled defines if Toggle is true or false
		Toggled bool
	}
	// Submit is the submit button
	Submit struct {
		// Text is the text of the Submit button
		Text string
		// Interaction is the action that will be taken when I click the trivial button
		Interaction bool
		// CloseFormOnSubmit defines if when I click Submit button it has to close the form
		CloseFormOnSubmit bool
	}
}

func (c CustomForm) CustomFormTitle() string {
	return c.Title
}
func (c CustomForm) CustomFormContent() string {
	return c.Content
}
func (c CustomForm) CustomFormTextBox() string {
	return c.TextBox
}
func (c CustomForm) CustomFormSliderText() string {
	return c.Slider.Text
}
func (c CustomForm) CustomFormSliderMax() int32 {
	return c.Slider.Max
}
func (c CustomForm) CustomFormSliderMin() int32 {
	return c.Slider.Min
}
func (c CustomForm) CustomFormDropdownText() string {
	return c.Dropdown.Text
}
func (c CustomForm) CustomFormDropdownArguments() []string {
	return c.Dropdown.Arguments
}
func (c CustomForm) CustomFormToggleText() string {
	return c.Toggle.Text
}
func (c CustomForm) CustomFormToggleToggled() bool {
	return c.Toggle.Toggled
}
func (c CustomForm) CustomFormSubmitText() string {
	return c.Submit.Text
}
func (c CustomForm) CustomFormSubmitInteraction() bool {
	return c.Submit.Interaction
}
func (c CustomForm) CustomFormSubmitCloseFormOnSubmit() bool {
	return c.Submit.CloseFormOnSubmit
}
