package filteroverduetasks

import (
	"testing"
	"time"
)

func TestFilterOverdueBasic(t *testing.T) {
	base := time.Date(2025, 11, 10, 12, 0, 0, 0, time.UTC)
	tasks := []Task{
		{Title: "past incomplete", Due: base.Add(-2 * time.Hour), Completed: false},
		{Title: "future incomplete", Due: base.Add(2 * time.Hour), Completed: false},
		{Title: "past complete", Due: base.Add(-3 * time.Hour), Completed: true},
	}

	overdue := FilterOverdue(tasks, base)
	if len(overdue) != 1 || overdue[0].Title != "past incomplete" {
		t.Fatalf("expected one overdue task 'past incomplete', got %#v", overdue)
	}
}

func TestFilterOverdueEmpty(t *testing.T) {
	base := time.Now()
	overdue := FilterOverdue(nil, base)
	if overdue != nil {
		t.Fatalf("expected nil for empty input, got %#v", overdue)
	}
}

func TestFilterOverdueNone(t *testing.T) {
	base := time.Date(2025, 11, 10, 12, 0, 0, 0, time.UTC)
	tasks := []Task{{Title: "future", Due: base.Add(1 * time.Minute), Completed: false}}
	overdue := FilterOverdue(tasks, base)
	if overdue != nil {
		t.Fatalf("expected nil slice, got %#v", overdue)
	}
}

func TestFilterOverdueBoundary(t *testing.T) {
	base := time.Date(2025, 11, 10, 12, 0, 0, 0, time.UTC)
	tasks := []Task{
		{Title: "exact boundary", Due: base, Completed: false},
		{Title: "one ns before", Due: base.Add(-1 * time.Nanosecond), Completed: false},
	}
	overdue := FilterOverdue(tasks, base)
	if len(overdue) != 1 || overdue[0].Title != "one ns before" {
		t.Fatalf("expected only 'one ns before', got %#v", overdue)
	}
}
