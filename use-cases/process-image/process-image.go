package process

func ProcessImage(matrix *[3][3]float64, original *[][][]float64) *[][][]float64 {
	height := len((*original))
	width := len((*original)[0])
	destiny := NewEmptyMatrix(height, width)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			originalPoint := [3]float64{float64(x), float64(y), 1}

			newX := int(matrix[0][0]*originalPoint[0] +
				matrix[0][1]*originalPoint[1] +
				matrix[0][2]*originalPoint[2])

			newY := int(matrix[1][0]*originalPoint[0] +
				matrix[1][1]*originalPoint[1] +
				matrix[1][2]*originalPoint[2])

			if newX >= 0 && newX < width && newY >= 0 && newY < height {
				(*destiny)[newY][newX] = (*original)[y][x]
			}
		}
	}
	return destiny
}
