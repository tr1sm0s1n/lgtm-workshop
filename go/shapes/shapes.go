package shapes

import "math"

type Square struct {
	Length float64
}
type Circle struct {
	Radius float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}
func (s Square) Perimeter() float64 {
	return s.Length * 4
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
