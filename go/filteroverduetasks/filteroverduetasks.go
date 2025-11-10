package filteroverduetasks

import "time"

// Task represents a unit of work with a due date and completion state.
// Title: short human readable label.
// Due: absolute time when the task is expected to be completed.
// Completed: whether the task has already been finished.
type Task struct {
	Title     string
	Due       time.Time
	Completed bool
}

// FilterOverdue returns all tasks whose due date is strictly before the provided
// 'today' instant and that are not completed. The input slice order is preserved
// among the returned tasks.
func FilterOverdue(tasks []Task, today time.Time) []Task {
	if len(tasks) == 0 {
		return nil
	}

	overdue := make([]Task, 0, len(tasks))
	for _, t := range tasks {
		if !t.Completed && t.Due.Before(today) {
			overdue = append(overdue, t)
		}
	}
	if len(overdue) == 0 {
		return nil
	}
	return overdue
}
