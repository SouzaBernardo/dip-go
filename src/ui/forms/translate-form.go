package form

import (
	"fmt"
	usecases "pdi/src/use-cases"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewTranslateForm(matrix [][][]int, onSubmit func([][][]int)) {

	deltaXEntry := widget.NewEntry()
	deltaYEntry := widget.NewEntry()

	deltaXEntry.Resize(fyne.NewSize(500, 40))
	deltaYEntry.Resize(fyne.NewSize(500, 40))

	items := []*widget.FormItem{
		{Text: "Delta X", Widget: deltaXEntry},
		{Text: "Delta Y", Widget: deltaYEntry},
	}

	NewForm(items, func() {
		deltaX, errX := strconv.Atoi(deltaXEntry.Text)
		deltaY, errY := strconv.Atoi(deltaYEntry.Text)
		if errX != nil || errY != nil {
			dialog.ShowError(fmt.Errorf("erro ao converter valores de entrada: %v, %v", errX, errY), fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		result := usecases.TranslateMatrix(matrix, deltaX, deltaY)
		onSubmit(result)
	})
}
