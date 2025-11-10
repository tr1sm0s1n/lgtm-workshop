package main

import "testing"

func TestTotal(t *testing.T) {
	order := Order{
		Items: []OrderItem{
			{"Book", 25.0, 2},   // 50
			{"Pen", 5.0, 3},     // 15
			{"Notebook", 10.0, 1}, // 10
		},
	}

	expectedTotal := 75.0
	if total := order.Total(); total != expectedTotal {
		t.Errorf("Expected total %.2f, got %.2f", expectedTotal, total)
	}
}

func TestEligibleForFreeShipping(t *testing.T) {
	tests := []struct {
		name     string
		order    Order
		expected bool
	}{
		{
			name: "Total less than 100",
			order: Order{
				Items: []OrderItem{
					{"Item1", 20.0, 3}, // 60
				},
			},
			expected: false,
		},
		{
			name: "Total equal to 100",
			order: Order{
				Items: []OrderItem{
					{"Item1", 50.0, 2}, // 100
				},
			},
			expected: true,
		},
		{
			name: "Total greater than 100",
			order: Order{
				Items: []OrderItem{
					{"Item1", 40.0, 3}, // 120
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := tt.order.EligibleForFreeShipping(); result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
