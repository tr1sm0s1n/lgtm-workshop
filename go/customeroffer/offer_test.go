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
        // Edge cases - boundary conditions
        {"Exactly 18 years old and 3 purchases", Customer{Age: 18, PurchaseCount: 3}, true},
        {"Exactly 18 years old but only 2 purchases", Customer{Age: 18, PurchaseCount: 2}, false},
        {"19 years old but only 2 purchases", Customer{Age: 19, PurchaseCount: 2}, false},
        {"17 years old with 3 purchases", Customer{Age: 17, PurchaseCount: 3}, false},
        {"Zero age and purchases", Customer{Age: 0, PurchaseCount: 0}, false},
        {"Very high values", Customer{Age: 100, PurchaseCount: 1000}, true},
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