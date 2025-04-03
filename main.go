package main

import (
	"pdi/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	image := ui.NewImage("assets/image.jpg")
	image2 := ui.NewImage("assets/image.jpg")

	imageContainer := ui.NewImageSection("Imagem Original", image)
	imageContainer2 := ui.NewImageSection("Imagem Alterada", image2)

	header := ui.NewHeader(image2, imageContainer2)

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
