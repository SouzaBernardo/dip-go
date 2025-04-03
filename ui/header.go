package ui

import (
	"pdi/convert"
	usecases "pdi/use-cases"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateContainer(result [][][]float64, c *fyne.Container) {
	newImage := convert.ConvertMatrixToImage(result)
	newCanvasImage := canvas.NewImageFromImage(newImage)
	newCanvasImage.FillMode = canvas.ImageFillOriginal
	imageSection := NewImageSection("Imagem Alterada", newCanvasImage)
	c.Objects = []fyne.CanvasObject{imageSection}
	c.Refresh()
}

func NewHeader(image *canvas.Image, c *fyne.Container) *fyne.Container {
	matrix := convert.ConvertImageToMatrix(image)

	btn1 := widget.NewButton("Transladar", func() {
		NewForm2Param("Delta X: ", "Delta Y: ", func(x float64, y float64) {
			result := usecases.TranslateMatrix(matrix, x, y)
			updateContainer(result, c)
		})
	})

	btn2 := widget.NewButton("Escala", func() {
		NewForm1Param("Valor para escalar: ", func(value float64) {
			result := usecases.ScaleMatrix(matrix, value)
			updateContainer(result, c)
		})
	})

	btn3 := widget.NewButton("Rotação", func() {
		NewForm1Param("Valor para rotação: ", func(value float64) {
			result := usecases.RotationMatrix(matrix, value)
			updateContainer(result, c)
		})
	})

	btn4 := widget.NewButton("Espelhamento", func() {
		result := usecases.MirrorMatrix(matrix)
		updateContainer(result, c)
	})

	btn5 := widget.NewButton("Brilho", func() {
		NewForm1Param("Valor para brilho: ", func(value float64) {
			result := usecases.ContrastMatrix(matrix, 1, value)
			updateContainer(result, c)
		})
	})

	btn6 := widget.NewButton("Contraste", func() {
		NewForm1Param("Valor para contraste: ", func(value float64) {
			result := usecases.ContrastMatrix(matrix, value, 1)
			updateContainer(result, c)
		})
	})

	btn7 := widget.NewButton("Gray Scale", func() {
		result := usecases.GrayScaleMatrix(matrix)
		updateContainer(result, c)
	})

	btn8 := widget.NewButton("Filtragem da média", func() {
		result := usecases.AverageFiltering(matrix)
		updateContainer(result, c)
	})

	header := container.NewHBox(btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
	return container.NewCenter(header)
}
