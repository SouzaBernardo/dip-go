package process

func NewEmptyMatrix(height int, width int) *[][][]float64 {

	NewEmptyMatrix := make([][][]float64, height)
	for i := range NewEmptyMatrix {
		NewEmptyMatrix[i] = make([][]float64, width)
		for j := range NewEmptyMatrix[i] {
			NewEmptyMatrix[i][j] = make([]float64, 3) // R, G, B
		}
	}

	return &NewEmptyMatrix
}
