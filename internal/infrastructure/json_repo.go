package infrastructure

import (
	"encoding/json"
	"os"

	"github.com/AndrivA89/fyne-todo-demo/internal/domain"
)

type JSONRepository struct {
	FilePath string
}

func NewJSONRepository(filePath string) *JSONRepository {
	return &JSONRepository{FilePath: filePath}
}

func (r *JSONRepository) Load() ([]domain.Task, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []domain.Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *JSONRepository) Save(tasks []domain.Task) error {
	file, err := os.Create(r.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tasks)
}
