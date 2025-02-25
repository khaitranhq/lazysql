package components

import "github.com/rivo/tview"

type EnterPasswordForm struct {
	*tview.Form
}

func NewEnterPasswordForm() *EnterPasswordForm {
	form := tview.NewForm()
	form.AddPasswordField("Password", "", 0, '*', nil)
	form.AddButton("Connect", nil)
	return &EnterPasswordForm{
		Form: form,
	}
}
