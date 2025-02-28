package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func NewImage3(path string, w float32, h float32) *canvas.Image {
	image := canvas.NewImageFromFile(path)
	image.Resize(fyne.NewSize(w, h))
	image.FillMode = canvas.ImageFillOriginal
	return image
}

func NewImage(path string) *canvas.Image {
	return NewImage3(path, 255, 255)
}
