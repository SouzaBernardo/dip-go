package model

var grayScaleWeights []float64 = []float64{0.5, 0.419, 0.081}

func (m *Matrix) GrayScale() {

	destiny := NewMatrix(m.Height, m.Width)
	destiny.ForEachPixel(func(x, y int) {
		rgb := (*m.Value)[y][x]
		grayValue := 0.0

		for i, value := range rgb {
			grayValue += value * grayScaleWeights[i]
		}

		gray := grayValue
		if gray > 255 {
			gray = 255
		}
		if gray < 0 {
			gray = 0
		}

		(*destiny.Value)[y][x] = []float64{gray, gray, gray}
	})
	m.Value = destiny.Value
	m.Height = destiny.Height
	m.Width = destiny.Width
}

func (m *Matrix) Contrast(contrast float64, glare float64) {
	destiny := NewMatrix(m.Height, m.Width)

	destiny.ForEachPixel(func(x, y int) {
		rgb := (*m.Value)[y][x]
		for i := 0; i < len(rgb); i++ {
			rgb[i] = contrast*float64(rgb[i]) + glare

			if rgb[i] > 255 {
				rgb[i] = 255
			}
			if rgb[i] < 0 {
				rgb[i] = 0
			}
		}

		(*destiny.Value)[y][x] = rgb
	})
	m.Value = destiny.Value
	m.Height = destiny.Height
	m.Width = destiny.Width
}

func (m *Matrix) Filtering() {
	destiny := NewMatrix(m.Height, m.Width)

	mask := [3][3]float64{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}

	for y := 1; y < destiny.Height-1; y++ {
		for x := 1; x < destiny.Width-1; x++ {
			z := [3]float64{0, 0, 0}

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					color := (*m.Value)[y+i-1][x+j-1]
					for k := 0; k < 3; k++ {
						z[k] += color[k] * mask[i][j]
					}
				}
			}

			for k := 0; k < 3; k++ {
				(*destiny.Value)[y][x][k] = z[k] / 16
			}
		}
	}
	m.Value = destiny.Value
	m.Height = destiny.Height
	m.Width = destiny.Width
}
