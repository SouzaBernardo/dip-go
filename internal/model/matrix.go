package model

type Matrix struct {
	Value  *[][][]float64
	Height int
	Width  int
}

func (m *Matrix) ForEachPixel(callback func(x, y int)) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			callback(x, y)
		}
	}
}

func (m *Matrix) ProcessMatrixRotation(matrix *[3][3]float64, original *[][][]float64) {
	height := len((*original))
	width := len((*original)[0])

	destiny := NewMatrix(height, width)

	halfX := float64(width) / 2.0
	halfY := float64(height) / 2.0

	destiny.ForEachPixel(func(x, y int) {
		tmpX := float64(x) - halfX
		tmpY := float64(y) - halfY

		newX := tmpX*(*matrix)[0][0] + tmpY*(*matrix)[0][1] + (*matrix)[0][2] + halfX
		newY := tmpX*(*matrix)[1][0] + tmpY*(*matrix)[1][1] + (*matrix)[1][2] + halfY

		if newX >= 0 && newX < float64(width) && newY >= 0 && newY < float64(height) {
			(*destiny.Value)[int(newY)][int(newX)] = (*original)[y][x]
		}
	})
	m.Value = destiny.Value
}

func (m *Matrix) ProcessMatrix(matrix *[3][3]float64) {

	destiny := NewMatrix(m.Height, m.Width)

	destiny.ForEachPixel(func(x, y int) {

		originalPoint := [3]float64{float64(x), float64(y), 1}

		newX := int(matrix[0][0]*originalPoint[0] +
			matrix[0][1]*originalPoint[1] +
			matrix[0][2]*originalPoint[2])

		newY := int(matrix[1][0]*originalPoint[0] +
			matrix[1][1]*originalPoint[1] +
			matrix[1][2]*originalPoint[2])

		if newX >= 0 && newX < destiny.Width && newY >= 0 && newY < destiny.Height {
			(*destiny.Value)[newY][newX] = (*m.Value)[y][x]
		}
	})

	m.Value = destiny.Value
}

func NewMatrix(height int, width int) *Matrix {
	value := make([][][]float64, height)
	for i := range value {
		value[i] = make([][]float64, width)
		for j := range value[i] {
			value[i][j] = make([]float64, 3)
		}
	}

	return &Matrix{
		Value:  &value,
		Height: height,
		Width:  width,
	}
}
