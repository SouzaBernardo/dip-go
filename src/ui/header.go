package ui

import (
	usecases "pdi/src/use-cases"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateContainer(result [][]int, c *fyne.Container) {
	newImage := usecases.ConvertMatrixToImage(result)
	newCanvasImage := canvas.NewImageFromImage(newImage)
	newCanvasImage.FillMode = canvas.ImageFillOriginal
	imageSection := NewImageSection("Imagem Alterada", newCanvasImage)
	c.Objects = []fyne.CanvasObject{imageSection}
	c.Refresh()
}

func NewHeader(image *canvas.Image, c *fyne.Container) *fyne.Container {
	matrix := usecases.ConvertImageToMatrix(image)

	btn1 := widget.NewButton("Inverter", func() {
		result := usecases.InvertImage(matrix)
		updateContainer(result, c)
	})
	btn2 := widget.NewButton("Transladar", func() {
		TranslateForm().Show()
	})
	btn3 := widget.NewButton("Escala", func() {})
	btn4 := widget.NewButton("Opção 4", func() {})
	btn5 := widget.NewButton("Opção 5", func() {})

	header := container.NewHBox(btn1, btn2, btn3, btn4, btn5)
	return container.NewCenter(header)
}
