package convert

import (
	"fyne.io/fyne/v2/canvas"
)

func ConvertImageToMatrix(img *canvas.Image) [][][]float64 {
	bounds := img.Image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	matrix := make([][][]float64, height)
	for y := 0; y < height; y++ {
		matrix[y] = make([][]float64, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.Image.At(x, y).RGBA()
			matrix[y][x] = []float64{float64(r >> 8), float64(g >> 8), float64(b >> 8)}
		}
	}

	return matrix
}
