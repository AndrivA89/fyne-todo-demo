package application

import (
	"fyne.io/fyne/v2"
	"github.com/google/uuid"

	"github.com/AndrivA89/fyne-todo-demo/internal/constants"
	"github.com/AndrivA89/fyne-todo-demo/internal/domain"
	"github.com/AndrivA89/fyne-todo-demo/internal/ports"
)

type TaskService struct {
	repo          ports.TaskRepository
	Tasks         []domain.Task
	TaskContainer *fyne.Container
	RefreshTasks  func()
}

func NewTaskService(
	repo ports.TaskRepository,
	taskContainer *fyne.Container,
) *TaskService {
	return &TaskService{
		repo:          repo,
		TaskContainer: taskContainer,
	}
}

func (s *TaskService) LoadTasks() ([]domain.Task, error) {
	return s.repo.Load()
}

func (s *TaskService) SaveTasks(tasks []domain.Task) error {
	return s.repo.Save(tasks)
}

func (s *TaskService) AddTask(tasks []domain.Task, text string) []domain.Task {
	newTask := domain.NewTask(text)
	return append(tasks, newTask)
}

func (s *TaskService) FilterTasks(tasks []domain.Task, filter string) []domain.Task {
	var filtered []domain.Task

	switch filter {
	case constants.FilterActive:
		for _, task := range tasks {
			if !task.Completed {
				filtered = append(filtered, task)
			}
		}
	case constants.FilterCompleted:
		for _, task := range tasks {
			if task.Completed {
				filtered = append(filtered, task)
			}
		}
	default:
		filtered = tasks
	}

	return filtered
}

func (s *TaskService) DeleteTaskByID(id uuid.UUID) {
	for i, task := range s.Tasks {
		if task.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			break
		}
	}
}

func (s *TaskService) EditTaskByID(id uuid.UUID, newText string) {
	for i, task := range s.Tasks {
		if task.ID == id {
			s.Tasks[i].SetText(newText)
			break
		}
	}
}

func (s *TaskService) MoveTaskByID(id uuid.UUID, direction string) {
	for i, task := range s.Tasks {
		if task.ID == id {
			if direction == constants.UpButton && i > 0 {
				s.Tasks[i], s.Tasks[i-1] = s.Tasks[i-1], s.Tasks[i]
			} else if direction == constants.DownButton && i < len(s.Tasks)-1 {
				s.Tasks[i], s.Tasks[i+1] = s.Tasks[i+1], s.Tasks[i]
			}
			break
		}
	}
}

func (s *TaskService) UpdateUI(tasks []domain.Task, createTaskRow func(task domain.Task) fyne.CanvasObject) {
	s.TaskContainer.Objects = nil
	for _, task := range tasks {
		s.TaskContainer.Add(createTaskRow(task))
	}
	s.TaskContainer.Refresh()
}
