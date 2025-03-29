package usecases

func TranslateMatrix(matrix [][][]int, deltaX int, deltaY int) [][][]int {

	height := len(matrix)
	width := len(matrix[0])

	translatedMatrix := make([][][]int, height)
	for i := range translatedMatrix {
		translatedMatrix[i] = make([][]int, width)
		for j := range translatedMatrix[i] {
			translatedMatrix[i][j] = make([]int, 3) // R, G, B
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newX := x + deltaX
			newY := y + deltaY
			if newX >= 0 && newX < width && newY >= 0 && newY < height {
				translatedMatrix[newY][newX] = matrix[y][x]
			}
		}
	}

	return translatedMatrix
}
