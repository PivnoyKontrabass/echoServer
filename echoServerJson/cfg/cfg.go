package cfg

import "time"

type TaskMetadata struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by" validate:"required"`
	Priority  int       `json:"priority" validate:"gte=0"`
	IsPrivate bool      `json:"is_private"`
}

// Подзадача (вложенная структура)
type Subtask struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// Основная структура задачи
type Task struct {
	ID          string       `json:"id"`
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"description"`
	Status      string       `json:"status" validate:"required,oneof=todo in_progress done"`
	Metadata    TaskMetadata `json:"metadata" validate:"required"`
	Subtasks    []Subtask    `json:"subtasks"`
}
