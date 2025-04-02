package form

import (
	"fmt"
	usecases "pdi/src/use-cases"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewScaleForm(matrix [][][]float64, onSubmit func([][][]float64)) {
	scale := widget.NewEntry()
	items := []*widget.FormItem{{Text: "Valor da Escala:", Widget: scale}}

	NewForm(items, func() {
		scaleInt, errX := strconv.ParseFloat(scale.Text, 64)
		if errX != nil {
			dialog.ShowError(fmt.Errorf("erro ao converter valores de entrada: %v", errX), fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		result := usecases.ScaleMatrix(matrix, scaleInt)
		onSubmit(result)
	})
}
