package employeeattendance

// Employee represents an employee's attendance record
type Employee struct {
	PresentDays  int
	WorkingDays  int
}

// AttendanceRate calculates the attendance percentage
func (e Employee) AttendanceRate() float64 {
	if e.WorkingDays == 0 {
		return 0.0
	}
	return (float64(e.PresentDays) / float64(e.WorkingDays)) * 100
}

// IsRegular returns true if attendance rate is at least 80%
func (e Employee) IsRegular() bool {
	return e.AttendanceRate() >= 80.0
}
