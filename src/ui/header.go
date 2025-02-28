package ui

import (
	usecases "pdi/src/use-cases"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewHeader(image [][]int) *fyne.Container {
	btn1 := widget.NewButton("Inverter",func() {usecases.InvertImage(image)} )
	btn2 := widget.NewButton("Opção 2", func() {})
	btn3 := widget.NewButton("Opção 3", func() {})
	btn4 := widget.NewButton("Opção 4", func() {})

	header := container.NewHBox(
		btn1, btn2, btn3, btn4,
	)
	return container.NewCenter(header)
}
