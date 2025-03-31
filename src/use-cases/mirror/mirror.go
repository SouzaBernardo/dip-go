package mirror

import "pdi/src/use-cases/process-image"

var counter int = 0

func MirrorMatrix(matrix [][][]int) [][][]int {

	pair := counter%2 == 0
	counter++

	if pair {
		transformationMatrix := &[3][3]int{
			{-1, 0, len(matrix[0])},
			{0, 1, 0},
			{0, 0, 1},
		}
		return *process.ProcessImage(transformationMatrix, &matrix)
	}
	transformationMatrix := &[3][3]int{
		{1, 0, 0},
		{0, -1, len(matrix)},
		{0, 0, 1},
	}
	return *process.ProcessImage(transformationMatrix, &matrix)
}
