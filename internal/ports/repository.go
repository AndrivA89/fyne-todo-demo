package ports

import (
	"github.com/AndrivA89/fyne-todo-demo/internal/domain"
)

type TaskRepository interface {
	Save(tasks []domain.Task) error
	Load() ([]domain.Task, error)
}
