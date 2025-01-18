package ui

import (
	"strings"

	"github.com/google/uuid"

	"github.com/AndrivA89/fyne-todo-demo/internal/application"
	"github.com/AndrivA89/fyne-todo-demo/internal/domain"
)

type TaskViewModel struct {
	Service *application.TaskService
}

func NewTaskViewModel(service *application.TaskService) *TaskViewModel {
	return &TaskViewModel{
		Service: service,
	}
}

func (vm *TaskViewModel) AddTask(text string) {
	if strings.TrimSpace(text) != "" {
		vm.Service.Tasks = vm.Service.AddTask(vm.Service.Tasks, text)
		_ = vm.Service.SaveTasks(vm.Service.Tasks)
	}
}

func (vm *TaskViewModel) EditTask(id uuid.UUID, newText string) {
	if strings.TrimSpace(newText) != "" {
		vm.Service.EditTaskByID(id, newText)
		_ = vm.Service.SaveTasks(vm.Service.Tasks)
	}
}

func (vm *TaskViewModel) DeleteTask(id uuid.UUID) {
	vm.Service.DeleteTaskByID(id)
	_ = vm.Service.SaveTasks(vm.Service.Tasks)
}

func (vm *TaskViewModel) ToggleTask(id uuid.UUID, completed bool) {
	for i := range vm.Service.Tasks {
		if vm.Service.Tasks[i].ID == id {
			vm.Service.Tasks[i].SetCompleted(completed)
			break
		}
	}
	_ = vm.Service.SaveTasks(vm.Service.Tasks)
}

func (vm *TaskViewModel) FilterTasks(filter string) []domain.Task {
	return vm.Service.FilterTasks(vm.Service.Tasks, filter)
}

func (vm *TaskViewModel) MoveTask(id uuid.UUID, direction string) {
	vm.Service.MoveTaskByID(id, direction)
	_ = vm.Service.SaveTasks(vm.Service.Tasks)
}
