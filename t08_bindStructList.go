package main

type t08_todo struct {
	Title    string
	Finished bool
}

func t08_bindStruct() {
	todos := []t08_todo{
		t08_todo{Title: "Title 1", Finished: false},
		t08_todo{Title: "Title 2", Finished: true},
	}

	// How do I bind a list of structs?
}
