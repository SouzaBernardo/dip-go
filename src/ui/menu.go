package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func NewMenu() {
	image := NewImage("src/assets/image.jpg")
	image2 := NewImage("src/assets/image.jpg")

	imageContainer := NewImageSection("Imagem Original", image)
	imageContainer2 := NewImageSection("Imagem Alterada", image2)

	header := NewHeader([][]int{})

	a := app.New()
	w := a.NewWindow("Trabalho")
	w.Resize(fyne.NewSize(550, 280))

	w.SetContent(
		container.NewBorder(
			header,
			nil,
			imageContainer,
			imageContainer2,
		),
	)
	w.ShowAndRun()
}
