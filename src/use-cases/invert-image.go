package usecases

func InvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			image[i][j] = 255 - image[i][j]
		}
	}
	return image
}
