package domain

import "github.com/google/uuid"

type Task struct {
	ID        uuid.UUID
	Text      string
	Completed bool
}

func (t *Task) SetCompleted(check bool) {
	t.Completed = check
}

func (t *Task) SetText(newText string) {
	t.Text = newText
}

func NewTask(text string) Task {
	return Task{
		ID:        uuid.New(),
		Text:      text,
		Completed: false,
	}
}
