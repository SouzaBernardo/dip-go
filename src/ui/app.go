package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func NewApp() {
	image := NewImage("src/assets/image.jpg")
	image2 := NewImage("src/assets/image.jpg")

	imageContainer := NewImageSection("Imagem Original", image)
	imageContainer2 := NewImageSection("Imagem Alterada", image2)

	header := NewHeader(image2, imageContainer2)

	a := app.New()
	w := a.NewWindow("Trabalho PDI")
	w.Resize(fyne.NewSize(750, 380))
	w.SetFixedSize(true)

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
