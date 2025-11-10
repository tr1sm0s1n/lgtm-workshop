package employeeattendance

import "testing"

func TestAttendanceRate(t *testing.T) {
	e := Employee{PresentDays: 24, WorkingDays: 30}
	got := e.AttendanceRate()
	want := 80.0

	if got != want {
		t.Errorf("AttendanceRate() = %v; want %v", got, want)
	}
}

func TestIsRegular(t *testing.T) {
	tests := []struct {
		name string
		e    Employee
		want bool
	}{
		{"Regular Employee", Employee{PresentDays: 28, WorkingDays: 30}, true},
		{"Irregular Employee", Employee{PresentDays: 10, WorkingDays: 30}, false},
		{"Zero Working Days", Employee{PresentDays: 0, WorkingDays: 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.IsRegular(); got != tt.want {
				t.Errorf("IsRegular() = %v; want %v", got, tt.want)
			}
		})
	}
}
