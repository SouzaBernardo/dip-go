package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewImage(path string) *canvas.Image {
	image := canvas.NewImageFromFile(path)
	image.Resize(fyne.NewSize(255, 255))
	image.FillMode = canvas.ImageFillOriginal
	return image
}
