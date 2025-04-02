package usecases

func AverageFiltering(matrix [][][]int) [][][]int {
    height := len(matrix)
    width := len(matrix[0])

    // Criar uma nova matriz para armazenar o resultado
    newMatrix := make([][][]int, height)
    for i := range newMatrix {
        newMatrix[i] = make([][]int, width)
        for j := range newMatrix[i] {
            newMatrix[i][j] = make([]int, 3) // R, G, B
        }
    }

    // Máscara do filtro da média (3x3 com todos os valores iguais a 1)
    mask := [3][3]int{
        {1, 1, 1},
        {1, 1, 1},
        {1, 1, 1},
    }

    // Aplicar o filtro da média
    for y := 1; y < height-1; y++ {
        for x := 1; x < width-1; x++ {
            sum := [3]int{0, 0, 0} // Soma para R, G e B

            // Percorrer a máscara 3x3
            for i := 0; i < 3; i++ {
                for j := 0; j < 3; j++ {
                    neighbor := matrix[y+i-1][x+j-1]
                    for k := 0; k < 3; k++ { // Para cada canal (R, G, B)
                        sum[k] += neighbor[k] * mask[i][j]
                    }
                }
            }

            // Calcular a média dividindo pela soma dos pesos da máscara (9)
            for k := 0; k < 3; k++ {
                newMatrix[y][x][k] = sum[k] / 9

                // Garantir que os valores estejam no intervalo [0, 255]
                if newMatrix[y][x][k] > 255 {
                    newMatrix[y][x][k] = 255
                }
                if newMatrix[y][x][k] < 0 {
                    newMatrix[y][x][k] = 0
                }
            }
        }
    }

    return newMatrix
}