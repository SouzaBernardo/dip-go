package model

import (
	"math"
)

var counter int = 0

func (m *Matrix) Translate(x float64, y float64) {
	matrix := &[3][3]float64{
		{1, 0, x},
		{0, 1, y},
		{0, 0, 1},
	}
	m.ProcessMatrix(matrix)
}

func (m *Matrix) Rotate(value float64) {
	radians := value * math.Pi / 180
	matrix := &[3][3]float64{
		{math.Cos(radians), -math.Sin(radians), 0},
		{math.Sin(radians), math.Cos(radians), 0},
		{0, 0, 1},
	}
	m.ProcessMatrixRotation(matrix, m.Value)
}

func (m *Matrix) Scale(value float64) {
	matrix := &[3][3]float64{
		{value, 0, 0},
		{0, value, 0},
		{0, 0, 1},
	}
	m.ProcessMatrix(matrix)
}

func (m *Matrix) Mirror() {
	pair := counter%2 == 0
	counter++

	matrix := &[3][3]float64{
		{1, 0, 0},
		{0, -1, float64(len(*m.Value))},
		{0, 0, 1},
	}

	if pair {
		matrix = &[3][3]float64{
			{-1, 0, float64(len((*m.Value)[0]))},
			{0, 1, 0},
			{0, 0, 1},
		}
	}

	m.ProcessMatrix(matrix)
}
