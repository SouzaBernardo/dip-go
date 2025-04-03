package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewForm1Param(inputText string, callback func(value float64)) {
	inputValue := widget.NewEntry()
	items := []*widget.FormItem{{Text: inputText, Widget: inputValue}}

	newForm(items, func() {
		valueParsed, errX := strconv.ParseFloat(inputValue.Text, 64)
		if errX != nil {
			dialog.ShowError(fmt.Errorf("erro ao converter valores de entrada: %v", errX), fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		callback(valueParsed)
	})
}

func NewForm2Param(inputText1 string, inputText2 string, callback func(value1 float64, value2 float64)) {
	inputValue1 := widget.NewEntry()
	inputValue2 := widget.NewEntry()
	items := []*widget.FormItem{
		{Text: inputText1, Widget: inputValue1},
		{Text: inputText2, Widget: inputValue2},
	}

	newForm(items, func() {
		valueParsed1, errX := strconv.ParseFloat(inputValue1.Text, 64)
		valueParsed2, errY := strconv.ParseFloat(inputValue2.Text, 64)
		if errX != nil || errY != nil {
			dialog.ShowError(fmt.Errorf("erro ao converter valores de entrada: %v %v", errX, errY), fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}
		callback(valueParsed1, valueParsed2)
	})
}

func newForm(items []*widget.FormItem, onSubmit func()) {
	dialog.NewForm("Aplicar", "OK", "Cancel", items, func(confirm bool) {
		if confirm {
			onSubmit()
		}
	}, fyne.CurrentApp().Driver().AllWindows()[0]).Show()
}
