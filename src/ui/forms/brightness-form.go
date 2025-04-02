package form

import (
	"fmt"
	usecases "pdi/src/use-cases"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewBrightnessForm(matrix [][][]float64, onSubmit func([][][]float64)) {
	brightness := widget.NewEntry()
	items := []*widget.FormItem{{Text: "Valor para brilho: ", Widget: brightness}}

	NewForm(items, func() {
		glare, errX := strconv.ParseFloat(brightness.Text, 64)
		if errX != nil {
			dialog.ShowError(fmt.Errorf("erro ao converter valores de entrada: %v", errX), fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		result := usecases.BirghtnessMatrix(matrix, glare)
		onSubmit(result)
	})

}
