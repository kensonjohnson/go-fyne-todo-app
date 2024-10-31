package main

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
)

type Todo struct {
	Description string
	Done        bool
}

func newTodo(description string) Todo {
	return Todo{Description: description}
}

func (t Todo) String() string {
	return fmt.Sprintf("%s - %t", t.Description, t.Done)
}

func newTodoFromDataItem(item binding.DataItem) Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(Todo)
}
