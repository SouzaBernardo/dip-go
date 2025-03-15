package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func TranslateForm() *dialog.FormDialog {

	deltaXEntry := widget.NewEntry()
	deltaYEntry := widget.NewEntry()

	items := []*widget.FormItem{
		{Text: "Delta X", Widget: deltaXEntry},
		{Text: "Delta Y", Widget: deltaYEntry},
	}
	// form := NewForms(items, deltaXEntry, deltaYEntry)
	deltaXEntry.Resize(fyne.NewSize(500, 40))
	deltaYEntry.Resize(fyne.NewSize(500, 40))

	return dialog.NewForm("Transladar", "OK", "Cancel", items, func(b bool) {}, fyne.CurrentApp().Driver().AllWindows()[0])
}
