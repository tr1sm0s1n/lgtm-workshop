package main

import "fmt"

// Define OrderItem struct
type OrderItem struct {
	Name  string
	Price float64
	Qty   int
}

// Define Order struct
type Order struct {
	Items []OrderItem
}

// Method to calculate total
func (o Order) Total() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.Price * float64(item.Qty)
	}
	return total
}

// Method to check free shipping eligibility
func (o Order) EligibleForFreeShipping() bool {
	return o.Total() >= 100
}

// Main function to test
func main() {
	order := Order{
		Items: []OrderItem{
			{"Laptop Bag", 40.5, 1},
			{"Mouse", 15.0, 2},
			{"Keyboard", 30.0, 1},
		},
	}

	fmt.Printf("Total Order Amount: $%.2f\n", order.Total())
	fmt.Printf("Eligible for Free Shipping: %v\n", order.EligibleForFreeShipping())
}
