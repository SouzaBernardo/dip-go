package usecases

import "pdi/src/use-cases/process-image"

var counter int = 0

func MirrorMatrix(matrix [][][]float64) [][][]float64 {

	pair := counter%2 == 0
	counter++

	transformationMatrix := &[3][3]float64{
		{1, 0, 0},
		{0, -1, float64(len(matrix))},
		{0, 0, 1},
	}

	if pair {
		transformationMatrix = &[3][3]float64{
			{-1, 0, float64(len(matrix[0]))},
			{0, 1, 0},
			{0, 0, 1},
		}
	}

	return *process.ProcessImageFloat(transformationMatrix, &matrix)
}
