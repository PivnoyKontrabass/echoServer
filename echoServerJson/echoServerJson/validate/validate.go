package validate

import (
	"echoServerJson/cfg"
	"strings"
	"unicode/utf8"
)

func ValidateTask(task *cfg.Task) map[string]string {
	errors := make(map[string]string)
	if strings.TrimSpace(task.Title) == "" {
		errors["title"] = "Title is required"
	}
	if utf8.RuneCountInString(task.Title) > 255 {
		errors["title"] = "Title too long (max 255 chars)"
	}
	if task.Status != "" {
		validStatus := map[string]bool{"todo": true, "in_progress": true, "done": true}
		if !validStatus[task.Status] {
			errors["status"] = "Invalid status. Use: todo, in_progress, done"
		}
	}
	if strings.TrimSpace(task.Metadata.CreatedBy) == "" {
		errors["metadata.created_by"] = "CreatedBy is required"
	}
	if task.Metadata.Priority <= 0 {
		errors["metadata.priority"] = "Priority must be greater or equal 0"
	}
	for _, subtask := range task.Subtasks {
		if strings.TrimSpace(subtask.Title) == "" {
			errors[subtask.Title] = "SubTask title is required"
		}
	}
	return errors
}
