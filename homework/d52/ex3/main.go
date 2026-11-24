package main

import "fmt"

const SantimPerBirr = 100

func main() {
	subtotalSantim := 68500
	deliverySantim := 3000

	serviceFeeSantim := int(float64(subtotalSantim) * 5 / 100)
	grandTotalSantim := subtotalSantim + serviceFeeSantim + deliverySantim

	fmt.Printf("Subtotal:    %.2f ETB\n", float64(subtotalSantim)/SantimPerBirr)
	fmt.Printf("Service fee: %.2f ETB\n", float64(serviceFeeSantim)/SantimPerBirr)
	fmt.Printf("Delivery:    %.2f ETB\n", float64(deliverySantim)/SantimPerBirr)
	fmt.Printf("Grand total: %.2f ETB\n", float64(grandTotalSantim)/SantimPerBirr)
}
