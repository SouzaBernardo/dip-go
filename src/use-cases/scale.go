package usecases

import "pdi/src/use-cases/process-image"

func ScaleMatrix(matrix [][][]int, scale float64) [][][]int {

	transformationMatrix := &[3][3]float64{
		{scale, 0, 0},
		{0, scale, 0},
		{0, 0, 1},
	}

	return *process.ProcessImageFloat(transformationMatrix, &matrix)
}
