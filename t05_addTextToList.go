package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type pinkEntryRenderer struct {
	fyne.WidgetRenderer
}

func (p *pinkEntryRenderer) BackgroundColor() color.Color {
	return color.RGBA{255, 20, 147, 255}
}

func t05_addTextToList() {
	myApp := app.New()

	win := myApp.NewWindow("TODOs")
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
