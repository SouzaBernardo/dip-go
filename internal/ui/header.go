package ui

import (
	"image"
	"image/color"
	"pdi/internal/model"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func convertMatrixToImage(matrix [][][]float64) *image.RGBA {
	height := len(matrix)
	width := len(matrix[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8(matrix[y][x][0])
			g := uint8(matrix[y][x][1])
			b := uint8(matrix[y][x][2])
			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	return img
}

func updateContainer(result [][][]float64, c *fyne.Container) {
	newImage := convertMatrixToImage(result)
	newCanvasImage := canvas.NewImageFromImage(newImage)
	newCanvasImage.FillMode = canvas.ImageFillOriginal
	imageSection := NewImageSection("Imagem Alterada", newCanvasImage)
	c.Objects = []fyne.CanvasObject{imageSection}
	c.Refresh()
}

func NewHeader(image *canvas.Image, c *fyne.Container) *fyne.Container {
	matrix := model.ConvertImageToMatrix(image)

	btn1 := widget.NewButton("Transladar", func() {
		NewForm2Param("Delta X: ", "Delta Y: ", func(x float64, y float64) {
			matrix.Translate(x, y)
			updateContainer(*matrix.Value, c)
		})
	})

	btn2 := widget.NewButton("Escala", func() {
		NewForm1Param("Valor para escalar: ", func(value float64) {
			matrix.Scale(value)
			updateContainer(*matrix.Value, c)
		})
	})

	btn3 := widget.NewButton("Rotação", func() {
		NewForm1Param("Valor para rotação: ", func(value float64) {
			matrix.Rotate(value)
			updateContainer(*matrix.Value, c)
		})
	})

	btn4 := widget.NewButton("Espelhamento", func() {
		matrix.Mirror()
		updateContainer(*matrix.Value, c)
	})

	btn5 := widget.NewButton("Brilho", func() {
		NewForm1Param("Valor para brilho: ", func(value float64) {
			matrix.Contrast(1, value)
			updateContainer(*matrix.Value, c)
		})
	})

	btn6 := widget.NewButton("Contraste", func() {
		NewForm1Param("Valor para contraste: ", func(value float64) {
			matrix.Contrast(value, 1)
			updateContainer(*matrix.Value, c)
		})
	})

	btn7 := widget.NewButton("Gray Scale", func() {
		matrix.GrayScale()
		updateContainer(*matrix.Value, c)
	})

	btn8 := widget.NewButton("Filtragem da média", func() {
		matrix.Filtering()
		updateContainer(*matrix.Value, c)
	})

	header := container.NewHBox(btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
	return container.NewCenter(header)
}
