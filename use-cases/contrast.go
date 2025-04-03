package usecases

import "pdi/use-cases/process-image"

func ContrastMatrix(matrix [][][]float64, contrast float64, glare float64) [][][]float64 {

	height := len(matrix)
	width := len(matrix[0])
	destiny := process.NewEmptyMatrix(height, width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgb := matrix[y][x]
			for i := 0; i < len(rgb); i++ {
				rgb[i] = contrast*float64(rgb[i]) + glare

				if rgb[i] > 255 {
					rgb[i] = 255
				}
				if rgb[i] < 0 {
					rgb[i] = 0
				}
			}

			(*destiny)[y][x] = rgb

		}
	}
	return *destiny
}
