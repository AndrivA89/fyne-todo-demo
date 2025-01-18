package main

import (
	"fyne.io/fyne/v2/app"

	"github.com/AndrivA89/fyne-todo-demo/internal/application"
	"github.com/AndrivA89/fyne-todo-demo/internal/infrastructure"
	"github.com/AndrivA89/fyne-todo-demo/internal/ui"
)

func main() {
	a := app.New()

	repo := infrastructure.NewJSONRepository("tasks.json")

	service := application.NewTaskService(repo, nil)
	service.Tasks, _ = service.LoadTasks()

	viewModel := ui.NewTaskViewModel(service)

	uiInstance := ui.NewFyneUI(a, viewModel)
	uiInstance.Run()
}
