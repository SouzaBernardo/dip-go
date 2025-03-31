package usecases

import (
	"pdi/src/use-cases/process-image"
)

func TranslateMatrix(matrix [][][]int, deltaX int, deltaY int) [][][]int {

	transformationMatrix := &[3][3]int{
		{1, 0, deltaX},
		{0, 1, deltaY},
		{0, 0, 1},
	}

	return *process.ProcessImage(transformationMatrix, &matrix)
}
