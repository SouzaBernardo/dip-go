package usecases

func BirghtnessMatrix(matrix [][][]int, glare float64) [][][]int {
	return ContrastMatrix(matrix, 1, glare)
}
