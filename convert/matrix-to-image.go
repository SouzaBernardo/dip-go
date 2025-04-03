package convert

import (
	"image"
	"image/color"
)

func ConvertMatrixToImage(matrix [][][]float64) *image.RGBA {
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
