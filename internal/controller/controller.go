package controller

import (
	"image"
	"image/color"
	"pdi/internal/model"

	"fyne.io/fyne/v2/canvas"
)

var matrix *model.Matrix = nil

func Init(image *canvas.Image) {
	matrix = convertImageToMatrix(image)
}

func Execute(service string, params ...interface{}) (*canvas.Image, bool) {
	if function, ok := Functions[service]; ok {
		result := function(params...)
		return convertMatrixToImage(result), true
	}
	return nil, false
}

func convertImageToMatrix(image *canvas.Image) *model.Matrix {
	bounds := image.Image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	matrix := make([][][]float64, height)
	for y := 0; y < height; y++ {
		matrix[y] = make([][]float64, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := image.Image.At(x, y).RGBA()
			matrix[y][x] = []float64{float64(r >> 8), float64(g >> 8), float64(b >> 8)}
		}
	}

	return &model.Matrix{
		Value:  &matrix,
		Height: height,
		Width:  width,
	}
}

func convertMatrixToImage(matrix *model.Matrix) *canvas.Image {
	image := image.NewRGBA(image.Rect(0, 0, matrix.Width, matrix.Height))
	matrix.ForEachPixel(func(x, y int) {
		r := uint8((*matrix.Value)[y][x][0])
		g := uint8((*matrix.Value)[y][x][1])
		b := uint8((*matrix.Value)[y][x][2])
		image.Set(x, y, color.RGBA{r, g, b, 255})
	})
	canvasImage := canvas.NewImageFromImage(image)
	canvasImage.FillMode = canvas.ImageFillOriginal
	return canvasImage
}
