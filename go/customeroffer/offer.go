package customeroffer

import "fmt"

type Customer struct {
	Age  int
	PurchaseCount int
}

func main() {

	fmt.Println("Enter the age")
	var age int
	fmt.Scanln(&age)
	fmt.Println("Enter the purchase count")
	var purchaseCount int
	fmt.Scanln(&purchaseCount)

	customer := Customer{Age: age, PurchaseCount: purchaseCount}
	offer := EligibleForOffer(customer)
	if offer {
		fmt.Println("Customer is eligible for the offer.")
	} else {
		fmt.Println("Customer is not eligible for the offer.")
	}
}

func EligibleForOffer(c Customer) bool {
	if c.Age >= 18 && c.PurchaseCount >= 3 {
		return true
	}
	return false
}