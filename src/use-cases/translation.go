package usecases

import (
	"pdi/src/use-cases/process-image"
)

func TranslateMatrix(matrix [][][]float64, deltaX float64, deltaY float64) [][][]float64 {

	transformationMatrix := &[3][3]float64{
		{1, 0, deltaX},
		{0, 1, deltaY},
		{0, 0, 1},
	}

	return *process.ProcessImageFloat(transformationMatrix, &matrix)
}
