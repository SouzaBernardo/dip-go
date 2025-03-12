package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func NewImageSection(message string, image *canvas.Image) *fyne.Container {
	fixedSize := fyne.NewSize(300, 25)
	rect := canvas.NewRectangle(color.Transparent)
	rect.SetMinSize(fixedSize)
	return container.NewVBox(
		container.NewCenter(canvas.NewText(message, color.White)),
		container.NewCenter(image),
		rect,
	)
}
