package form

import (
	"fmt"
	usecases "pdi/src/use-cases"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewRotationForm(matrix [][][]float64, onSubmit func([][][]float64)) {
	rotation := widget.NewEntry()
	items := []*widget.FormItem{{Text: "Valor para rotação: ", Widget: rotation}}

	NewForm(items, func() {
		rotationInt, errX := strconv.ParseFloat(rotation.Text, 64)
		if errX != nil {
			dialog.ShowError(fmt.Errorf("erro ao converter valores de entrada: %v", errX), fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		result := usecases.RotationMatrix(matrix, rotationInt)
		onSubmit(result)
	})
}
