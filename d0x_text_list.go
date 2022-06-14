package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// https://developer.fyne.io/widget/list
func textList() {
	a := app.New()
	w := a.NewWindow("TrinityList")
	w.Resize(fyne.NewSize(600, 800))

	var data = []string{"a", "string", "list"}

	list := widget.NewList(
		// length
		func() int {
			return len(data)
		},
		// createItem
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		//updateItem
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(data[lii])
		},
	)
	w.SetContent(list)
	w.ShowAndRun()
}
