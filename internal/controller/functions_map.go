package controller

import "pdi/internal/model"

var Functions = map[string]func(...interface{}) *model.Matrix{
	"translate": func(params ...interface{}) *model.Matrix {
		x := params[0].(float64)
		y := params[1].(float64)
		matrix.Translate(x, y)
		return matrix
	},
	"scale": func(params ...interface{}) *model.Matrix {
		value := params[0].(float64)
		matrix.Scale(value)
		return matrix
	},
	"rotate": func(params ...interface{}) *model.Matrix {
		value := params[0].(float64)
		matrix.Rotate(value)
		return matrix
	},
	"mirror": func(params ...interface{}) *model.Matrix {
		matrix.Mirror()
		return matrix
	},
	"contrast": func(params ...interface{}) *model.Matrix {
		contrast := params[0].(float64)
		matrix.Contrast(contrast, 1)
		return matrix
	},
	"filtering": func(params ...interface{}) *model.Matrix {
		matrix.Filtering()
		return matrix
	},
	"brightness": func(params ...interface{}) *model.Matrix {
		glare := params[0].(float64)
		matrix.Contrast(1, glare)
		return matrix
	},
	"grayscale": func(params ...interface{}) *model.Matrix {
		matrix.GrayScale()
		return matrix
	},
}
