package application

import (
	"testing"

	"github.com/google/uuid"

	"github.com/AndrivA89/fyne-todo-demo/internal/domain"
)

type mockTaskRepository struct {
	tasks []domain.Task
}

func (m *mockTaskRepository) Load() ([]domain.Task, error) {
	return m.tasks, nil
}

func (m *mockTaskRepository) Save(tasks []domain.Task) error {
	m.tasks = tasks
	return nil
}

func TestTaskService_AddTask(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskService(mockRepo, nil)

	var tasks []domain.Task
	tasks = service.AddTask(tasks, "Test Task")

	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Text != "Test Task" {
		t.Fatalf("expected task text to be 'Test Task', got '%s'", tasks[0].Text)
	}

	if tasks[0].ID == uuid.Nil {
		t.Fatalf("expected task to have a valid UUID, got nil")
	}
}

func TestTaskService_FilterTasks(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskService(mockRepo, nil)

	tasks := []domain.Task{
		{ID: uuid.New(), Text: "Task 1", Completed: false},
		{ID: uuid.New(), Text: "Task 2", Completed: true},
		{ID: uuid.New(), Text: "Task 3", Completed: false},
	}

	activeTasks := service.FilterTasks(tasks, "Active")
	if len(activeTasks) != 2 {
		t.Fatalf("expected 2 active tasks, got %d", len(activeTasks))
	}

	completedTasks := service.FilterTasks(tasks, "Completed")
	if len(completedTasks) != 1 {
		t.Fatalf("expected 1 completed task, got %d", len(completedTasks))
	}

	allTasks := service.FilterTasks(tasks, "All")
	if len(allTasks) != 3 {
		t.Fatalf("expected 3 tasks, got %d", len(allTasks))
	}
}

func TestTaskService_EditTaskByID(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskService(mockRepo, nil)

	taskID := uuid.New()
	tasks := []domain.Task{
		{ID: taskID, Text: "Old Task"},
	}

	service.Tasks = tasks
	service.EditTaskByID(taskID, "New Task")

	if service.Tasks[0].Text != "New Task" {
		t.Fatalf("expected task text to be 'New Task', got '%s'", service.Tasks[0].Text)
	}
}

func TestTaskService_DeleteTaskByID(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskService(mockRepo, nil)

	taskID1 := uuid.New()
	taskID2 := uuid.New()

	tasks := []domain.Task{
		{ID: taskID1, Text: "Task 1"},
		{ID: taskID2, Text: "Task 2"},
	}

	service.Tasks = tasks
	service.DeleteTaskByID(taskID1)

	if len(service.Tasks) != 1 {
		t.Fatalf("expected 1 remaining task, got %d", len(service.Tasks))
	}

	if service.Tasks[0].ID != taskID2 {
		t.Fatalf("expected remaining task ID to be '%s', got '%s'", taskID2, service.Tasks[0].ID)
	}
}

func TestTaskService_MoveTaskByID(t *testing.T) {
	mockRepo := &mockTaskRepository{}
	service := NewTaskService(mockRepo, nil)

	taskID1 := uuid.New()
	taskID2 := uuid.New()
	taskID3 := uuid.New()

	tasks := []domain.Task{
		{ID: taskID1, Text: "Task 1"},
		{ID: taskID2, Text: "Task 2"},
		{ID: taskID3, Text: "Task 3"},
	}

	service.Tasks = tasks
	service.MoveTaskByID(taskID2, "up")

	if service.Tasks[0].ID != taskID2 {
		t.Fatalf("expected first task to be '%s', got '%s'", taskID2, service.Tasks[0].ID)
	}

	service.MoveTaskByID(taskID2, "down")

	if service.Tasks[1].ID != taskID2 {
		t.Fatalf("expected second task to be '%s', got '%s'", taskID2, service.Tasks[1].ID)
	}
}
