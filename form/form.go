package form

type Form struct {
	SimpleForm SimpleForm
	CustomForm CustomForm
	ModalForm  ModalForm
}

func SendForm(form Form) Form {
	return form
}
func SendSimpleForm(form SimpleForm) SimpleForm {
	return form
}
func SendModalForm(form ModalForm) ModalForm {
	return form
}
func SendCustomForm(form CustomForm) CustomForm {
	return form
}
