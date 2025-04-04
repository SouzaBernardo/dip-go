package model

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
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	destiny.ForEachPixel(func(x, y int) {
		sum := [3]float64{0, 0, 0}

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				neighbor := (*m.Value)[y+i-1][x+j-1]
				for k := 0; k < 3; k++ {
					sum[k] += neighbor[k] * mask[i][j]
				}
			}
		}

		for k := 0; k < 3; k++ {
			(*destiny.Value)[y][x][k] = sum[k] / 9

			if (*destiny.Value)[y][x][k] > 255 {
				(*destiny.Value)[y][x][k] = 255
			}
			if (*destiny.Value)[y][x][k] < 0 {
				(*destiny.Value)[y][x][k] = 0
			}
		}
	})
	m.Value = destiny.Value
	m.Height = destiny.Height
	m.Width = destiny.Width
}
