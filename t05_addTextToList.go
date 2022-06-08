package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func t05_addTextToList() {
	myApp := app.NewWithID("bbkane.com/TODOs")

	win := myApp.NewWindow("TODOs")
	win.Resize(fyne.NewSize(600, 600))

	todoInputData := binding.BindString(ptr("hi"))

	todosListData := binding.NewStringList()

	list := widget.NewListWithData(
		todosListData,
		func() fyne.CanvasObject {
			btn_color := canvas.NewRectangle(
				color.NRGBA{R: 0, G: 255, B: 0, A: 255})
			btn_container := container.New(
				layout.NewMaxLayout(),
				btn_color,
				widget.NewLabel(""), // second object in container
			)
			return btn_container
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			// The second item is the label (see add func above)
			o.(*fyne.Container).Objects[1].(*widget.Label).Bind(i.(binding.String))
		},
	)

	addToList := func() {
		newItem, err := todoInputData.Get()
		if err != nil {
			fyne.LogError("Can't get todoInput", err)
		}
		todosListData.Append(newItem)
	}
	addBtn := widget.NewButton("Add", addToList)

	win.SetContent(
		container.NewBorder(
			nil,
			container.New(
				layout.NewFormLayout(),
				addBtn,
				func() *widget.Entry {
					wgt := widget.NewEntryWithData(todoInputData)
					wgt.OnSubmitted = func(s string) { addToList() }
					return wgt
				}(),
			),
			nil,
			nil,
			list,
		),
	)
	win.ShowAndRun()
}
