package form

import (
	"fmt"
	usecases "pdi/src/use-cases"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewContrastForm(matrix [][][]float64, onSubmit func([][][]float64)) {
	contrast := widget.NewEntry()
	items := []*widget.FormItem{{Text: "Valor para o contraste:", Widget: contrast}}

	NewForm(items, func() {
		contrastFloatValue, errX := strconv.ParseFloat(contrast.Text, 64)
		if errX != nil {
			dialog.ShowError(fmt.Errorf("erro ao converter valores de entrada: %v", errX), fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		result := usecases.ContrastMatrix(matrix, contrastFloatValue, 1)
		onSubmit(result)
	})

}
