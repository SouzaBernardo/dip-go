package usecases

import (
	"image"
	"image/color"
)

func ConvertMatrixToImage(matrix [][]int) *image.RGBA {
    height := len(matrix)
    width := len(matrix[0])
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            gray := uint8(matrix[y][x])
            img.Set(x, y, color.RGBA{gray, gray, gray, 255})
        }
    }

    return img
}