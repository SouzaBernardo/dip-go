package usecases

import (
	"math"
	"pdi/src/use-cases/process-image"
)

func RotationMatrix(matrix [][][]float64, rotation float64) [][][]float64 {
	radians := rotation * math.Pi / 180
	newMatrix := &[3][3]float64{
		{math.Cos(radians), -math.Sin(radians), 0},
		{math.Sin(radians), math.Cos(radians), 0},
		{0, 0, 1},
	}
	return *processMatrix(newMatrix, &matrix)

}

func processMatrix(matrix *[3][3]float64, original *[][][]float64) *[][][]float64 {

	height := len((*original))
	width := len((*original)[0])

	destiny := process.NewEmptyMatrix(height, width)

	halfX := float64(width) / 2.0
	halfY := float64(height) / 2.0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			tmpX := float64(x) - halfX
			tmpY := float64(y) - halfY

			newX := tmpX*(*matrix)[0][0] + tmpY*(*matrix)[0][1] + (*matrix)[0][2] + halfX
			newY := tmpX*(*matrix)[1][0] + tmpY*(*matrix)[1][1] + (*matrix)[1][2] + halfY

			if newX >= 0 && newX < float64(width) && newY >= 0 && newY < float64(height) {
				(*destiny)[int(newY)][int(newX)] = (*original)[y][x]
			}
		}
	}

	return destiny
}
