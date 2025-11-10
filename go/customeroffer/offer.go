package customeroffer

type Customer struct {
	Age  int
	PurchaseCount int
}


func EligibleForOffer(c Customer) bool {
	if c.Age >= 18 && c.PurchaseCount >= 3 {
		return true
	}
	return false
}