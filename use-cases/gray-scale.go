package usecases

import "pdi/use-cases/process-image"

var grayScaleWeights []float64 = []float64{0.5, 0.419, 0.081}

func GrayScaleMatrix(matrix [][][]float64) [][][]float64 {

	height := len((matrix))
	width := len(matrix[0])

	newMatrix := process.NewEmptyMatrix(height, width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgb := matrix[y][x]
			grayValue := 0.0

			for i, value := range rgb {
				grayValue += value * grayScaleWeights[i]
			}

			gray := grayValue
			if gray > 255 {
				gray = 255
			}
			if gray < 0 {
				gray = 0
			}

			(*newMatrix)[y][x] = []float64{gray, gray, gray}
		}
	}

	return *newMatrix
}
