package ui

import (
	form "pdi/src/ui/forms"
	usecases "pdi/src/use-cases"
	"pdi/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateContainer(result [][][]int, c *fyne.Container) {
	newImage := utils.ConvertMatrixToImage(result)
	newCanvasImage := canvas.NewImageFromImage(newImage)
	newCanvasImage.FillMode = canvas.ImageFillOriginal
	imageSection := NewImageSection("Imagem Alterada", newCanvasImage)
	c.Objects = []fyne.CanvasObject{imageSection}
	c.Refresh()
}

func NewHeader(image *canvas.Image, c *fyne.Container) *fyne.Container {
	matrix := utils.ConvertImageToMatrix(image)

	btn1 := widget.NewButton("Transladar", func() {
		form.TranslateForm(matrix, func(result [][][]int) { updateContainer(result, c) })
	})

	btn2 := widget.NewButton("Escala", func() {
		form.ScaleForm(matrix, func(result [][][]int) { updateContainer(result, c) })
	})

	btn3 := widget.NewButton("Rotação", func() {
		form.RotationForm(matrix, func(result [][][]int) { updateContainer(result, c) })
	})

	btn4 := widget.NewButton("Espelhamento", func() {
		result := usecases.MirrorMatrix(matrix)
		updateContainer(result, c)
	})

	btn5 := widget.NewButton("Brilho", func() {
		result := usecases.LightMatrix(matrix)
		updateContainer(result, c)
	})

	btn6 := widget.NewButton("Contraste", func() {
		form.NewContrastForm(matrix, func(result [][][]int) { updateContainer(result, c) })
	})

	btn7 := widget.NewButton("Gray Scale", func() {
		result := usecases.GrayScaleMatrix(matrix)
		updateContainer(result, c)
	})

	btn8 := widget.NewButton("Filtragem da média", func() {
		// result := usecases.MirorMatrix(matrix)
		// updateContainer(result, c)
	})

	header := container.NewHBox(btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
	return container.NewCenter(header)
}
