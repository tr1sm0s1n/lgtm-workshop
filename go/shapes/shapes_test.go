package shapes

import (
	"math"
	"testing"
)

func TestSquare(t *testing.T) {
	s := Square{Length: 5}

	if s.Area() != 25 {
		t.Fatalf("expected square area: %d, actual square area: %f", 25, s.Area())
	}

	if s.Perimeter() != 20 {
		t.Fatalf("expected square perimeter: %d, actual square perimeter: %f", 20, s.Perimeter())
	}
}

func TestCircle(t *testing.T) {
	c := Circle{Radius: 3}

	if c.Area() != 9*math.Pi {
		t.Fatalf("expected circle area: %f, actual circle area: %f", 9*math.Pi, c.Area())
	}

	if c.Perimeter() != 6*math.Pi {
		t.Fatalf("expected circle perimeter: %f, actual circle perimeter: %f", 6*math.Pi, c.Perimeter())
	}
}
