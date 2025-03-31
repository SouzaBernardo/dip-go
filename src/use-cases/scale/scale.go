package usecases

import "pdi/src/use-cases/process-image"

// todo: Ajustar para float
func ScaleMatrix(matrix [][][]int, scale int) [][][]int {

	transformationMatrix := &[3][3]int{
		{scale, 0, 0},
		{0, scale, 0},
		{0, 0, 1},
	}

	return *process.ProcessImage(transformationMatrix, &matrix)
}
