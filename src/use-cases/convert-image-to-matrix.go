package usecases

import "fyne.io/fyne/v2/canvas"

func ConvertImageToMatrix(img *canvas.Image) [][]int{
	bounds := img.Image.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	matrix := make([][]int, height)
	for y := 0; y < height; y++ {
		matrix[y] = make([]int, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.Image.At(x, y).RGBA() 
			gray := (r + g + b) / 3           
			matrix[y][x] = int(gray >> 8)     
		}
	}

	return matrix

}
