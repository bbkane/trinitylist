package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func t05_addTextToList() {
	myApp := app.New()

	win := myApp.NewWindow("List Data")
	win.Resize(fyne.NewSize(600, 600))

	todoInputData := binding.BindString(ptr("hi"))

	todosListData := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)
	list := widget.NewListWithData(
		todosListData,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	addBtn := widget.NewButton("Add", func() {
		newItem, err := todoInputData.Get()
		if err != nil {
			fyne.LogError("Can't get todoInput", err)
		}
		todosListData.Append(newItem)
	})

	win.SetContent(
		container.NewBorder(
			nil,
			container.New(
				layout.NewFormLayout(),
				addBtn,
				// widget.NewEntryWithData(
				// 	todoInputData,
				// ),
			),
			nil,
			nil,
			list,
		),
	)
	win.ShowAndRun()
}
