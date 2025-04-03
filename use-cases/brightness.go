package usecases

func BirghtnessMatrix(matrix [][][]float64, glare float64) [][][]float64 {
	return ContrastMatrix(matrix, 1, glare)
}
