package ui

import (
	"pdi/internal/controller"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateContainer(result *canvas.Image, c *fyne.Container) {
	imageSection := NewImageSection("Imagem Alterada", result)
	c.Objects = []fyne.CanvasObject{imageSection}
	c.Refresh()
}

func NewHeader(image *canvas.Image, c *fyne.Container) *fyne.Container {
	controller.Init(image)

	btn1 := widget.NewButton("Transladar", func() {
		NewForm2Param("Delta X: ", "Delta Y: ", func(x float64, y float64) {
			content, ok := controller.Execute("translate", x, y)
			if ok {
				updateContainer(content, c)
			}
		})
	})

	btn2 := widget.NewButton("Escala", func() {
		NewForm1Param("Valor para escalar: ", func(value float64) {
			content, ok := controller.Execute("scale", value)
			if ok {
				updateContainer(content, c)
			}
		})
	})

	btn3 := widget.NewButton("Rotação", func() {
		NewForm1Param("Valor para rotação: ", func(value float64) {
			content, ok := controller.Execute("rotate", value)
			if ok {
				updateContainer(content, c)
			}
		})
	})

	btn4 := widget.NewButton("Espelhamento", func() {
		content, ok := controller.Execute("mirror")
		if ok {
			updateContainer(content, c)
		}
	})

	btn5 := widget.NewButton("Brilho", func() {
		NewForm1Param("Valor para brilho: ", func(value float64) {
			content, ok := controller.Execute("brightness", value)
			if ok {
				updateContainer(content, c)
			}
		})
	})

	btn6 := widget.NewButton("Contraste", func() {
		NewForm1Param("Valor para contraste: ", func(value float64) {
			content, ok := controller.Execute("contrast", value)
			if ok {
				updateContainer(content, c)
			}
		})
	})

	btn7 := widget.NewButton("Gray Scale", func() {
		// matrix.GrayScale()
		// updateContainer(*matrix.Value, c)
		content, ok := controller.Execute("grayscale")
		if ok {
			updateContainer(content, c)
		}
	})

	btn8 := widget.NewButton("Filtragem da média", func() {
		content, ok := controller.Execute("filtering")
		if ok {
			updateContainer(content, c)
		}
	})

	header := container.NewHBox(btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8)
	return container.NewCenter(header)
}
