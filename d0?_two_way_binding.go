package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func bindingSimpleWidgets() {
	myApp := app.New()
	w := myApp.NewWindow("Two Way")

	str := binding.NewString()
	str.Set("Hi!")

	w.SetContent(container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	))

	w.ShowAndRun()
}
