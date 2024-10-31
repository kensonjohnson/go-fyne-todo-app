package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var sampleData = []Todo{
	newTodo("Some Stuff"),
	newTodo("Some more things"),
	newTodo("Yet another thing"),
}

func main() {
	// Create an app and window
	a := app.New()
	w := a.NewWindow("Todo App")

	// Customize window
	w.Resize(fyne.NewSize(300, 400))

	// Create data to bind to
	todos := binding.NewUntypedList()
	for _, todo := range sampleData {
		todos.Append(todo)
	}

	// Start adding widgets
	newTodoInput := widget.NewEntry()
	newTodoInput.PlaceHolder = "New Todo Description..."

	addButton := widget.NewButton("Add", func() {
		todos.Append(newTodo(newTodoInput.Text))
		newTodoInput.SetText("")
	})
	addButton.Disable()

	newTodoInput.OnChanged = func(s string) {
		addButton.Disable()

		if len(s) > 3 {
			addButton.Enable()
		}
	}

	w.SetContent(container.NewBorder(
		nil,
		container.NewBorder(
			nil,          // Top
			nil,          // Bottom
			nil,          // Left
			addButton,    // Right
			newTodoInput, // Remaining space
		),
		nil,
		nil,
		createTodoDataList(todos),
	))

	w.ShowAndRun()
	postCleanup()
}

func createTodoDataList(todos binding.UntypedList) *widget.List {
	return widget.NewListWithData(
		// The data list
		todos,
		// How to layout each widget
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil, nil,
				widget.NewCheck("", func(b bool) {}),
				widget.NewLabel(""),
			)
		},
		// How to render each component
		func(di binding.DataItem, o fyne.CanvasObject) {
			cont := o.(*fyne.Container)

			label := cont.Objects[0].(*widget.Label)
			check := cont.Objects[1].(*widget.Check)

			todo := newTodoFromDataItem(di)

			label.SetText(todo.Description)
			check.SetChecked(todo.Done)
		},
	)
}

func postCleanup() {
	fmt.Println("Exiting")
}
