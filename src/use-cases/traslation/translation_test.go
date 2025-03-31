package usecases

import (
	"reflect"
	"testing"
)

func TestTranslateMatrix(t *testing.T) {

	matrix := [][][]int{
		{{255, 0, 0}, {0, 255, 0}, {0, 0, 255}},
		{{255, 255, 0}, {0, 255, 255}, {255, 0, 255}},
		{{0, 0, 0}, {128, 128, 128}, {255, 255, 255}},
	}

	deltaX := 1
	deltaY := 1

	expected := [][][]int{
		{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		{{0, 0, 0}, {255, 0, 0}, {0, 255, 0}},
		{{0, 0, 0}, {255, 255, 0}, {0, 255, 255}},
	}

	result := TranslateMatrix(matrix, deltaX, deltaY)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Resultado incorreto. Esperado: %v, Obtido: %v", expected, result)
	}
}
