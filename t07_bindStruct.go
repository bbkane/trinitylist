package main

import (
	"log"

	"fyne.io/fyne/v2/data/binding"
)

type t07_task struct {
	Title    string
	Finished bool
}

func t07_bindStruct() {
	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)

	myInt := 5
	boundInt := binding.BindInt(&myInt)
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)

	t := t07_task{Finished: false, Title: "my title"}
	boundTask := binding.BindStruct(&t)

	{
		t, _ := boundTask.GetItem("Title")
		log.Printf("%#v\n", t)
	}

	{
		t, _ := boundTask.GetValue("Title")
		log.Printf("%#v\n", t)
	}

	t.Title = "My other title"

	{
		t, _ := boundTask.GetValue("Title")
		log.Printf("%#v\n", t)
	}
}
