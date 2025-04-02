package usecases

import "pdi/src/use-cases/process-image"

var grayScaleWeights []float64 = []float64{0.5, 0.419, 0.081}

func GrayScaleMatrix(matrix [][][]float64) [][][]int {

	height := len((matrix))
	width := len(matrix[0])

	newMatrix := process.NewEmptyMatrix(height, width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgb := matrix[y][x]
			grayValue := 0.0

			// Calcular o valor em escala de cinza usando os pesos
			for i, value := range rgb {
				grayValue += float64(value) * grayScaleWeights[i]
			}

			// Arredondar o valor e garantir que esteja no intervalo [0, 255]
			gray := int(grayValue)
			if gray > 255 {
				gray = 255
			}
			if gray < 0 {
				gray = 0
			}

			// Definir o mesmo valor para R, G e B na matriz de destino
			(*newMatrix)[y][x] = []int{gray, gray, gray}
		}
	}

	return *newMatrix
}
