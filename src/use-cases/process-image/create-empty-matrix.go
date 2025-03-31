package process

func NewEmptyMatrix(height int, width int) *[][][]int {

	NewEmptyMatrix := make([][][]int, height)
	for i := range NewEmptyMatrix {
		NewEmptyMatrix[i] = make([][]int, width)
		for j := range NewEmptyMatrix[i] {
			NewEmptyMatrix[i][j] = make([]int, 3) // R, G, B
		}
	}

	return &NewEmptyMatrix
}
