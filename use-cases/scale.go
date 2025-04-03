package usecases

import "pdi/use-cases/process-image"

func ScaleMatrix(matrix [][][]float64, scale float64) [][][]float64 {

	transformationMatrix := &[3][3]float64{
		{scale, 0, 0},
		{0, scale, 0},
		{0, 0, 1},
	}

	return *process.ProcessImage(transformationMatrix, &matrix)
}
