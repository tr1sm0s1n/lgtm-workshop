package customeroffer

import "testing"

func TestEligibleForOffer(t *testing.T) {
	tests := []struct {
		name     string
		customer Customer
		expected bool
	}{
		{"Eligible customer", Customer{Age: 25, PurchaseCount: 5}, true},
		{"Underage customer", Customer{Age: 16, PurchaseCount: 5}, false},
		{"Insufficient purchases", Customer{Age: 30, PurchaseCount: 2}, false},
		{"Underage and insufficient purchases", Customer{Age: 17, PurchaseCount: 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EligibleForOffer(tt.customer)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
