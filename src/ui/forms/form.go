package form

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewForm(items []*widget.FormItem, onSubmit func()) {
	dialog.NewForm("Transladar", "OK", "Cancel", items, func(confirm bool) {
		if confirm {
			onSubmit()
		}
	}, fyne.CurrentApp().Driver().AllWindows()[0]).Show()
}
