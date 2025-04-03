package usecases

func AverageFiltering(matrix [][][]float64) [][][]float64 {
	height := len(matrix)
	width := len(matrix[0])

	newMatrix := make([][][]float64, height)
	for i := range newMatrix {
		newMatrix[i] = make([][]float64, width)
		for j := range newMatrix[i] {
			newMatrix[i][j] = make([]float64, 3)
		}
	}

	mask := [3][3]float64{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			sum := [3]float64{0, 0, 0}

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					neighbor := matrix[y+i-1][x+j-1]
					for k := 0; k < 3; k++ {
						sum[k] += neighbor[k] * mask[i][j]
					}
				}
			}

			for k := 0; k < 3; k++ {
				newMatrix[y][x][k] = sum[k] / 9

				if newMatrix[y][x][k] > 255 {
					newMatrix[y][x][k] = 255
				}
				if newMatrix[y][x][k] < 0 {
					newMatrix[y][x][k] = 0
				}
			}
		}
	}

	return newMatrix
}
