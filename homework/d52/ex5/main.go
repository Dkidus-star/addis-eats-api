package main

import "fmt"

const (
	SantimPerBirr         = 100
	FreeDeliveryThreshold = 50000
)

func main() {
	areas := []string{"Bole", "Kazanchis", "Piassa", "Megenagna", "Ayat"}
	subtotals := []int{10000, 30000, 60000}

	fmt.Printf("%-12s | %-12s | %-12s | %-12s\n", "Area", "Subtotal", "Delivery Fee", "Grand Total")
	fmt.Println("---------------------------------------------------------")

	for _, area := range areas {
		for _, subtotal := range subtotals {
			var fee int
			switch area {
			case "Bole":
				fee = 3000
			case "Kazanchis":
				fee = 3500
			case "Piassa":
				fee = 4000
			case "Megenagna":
				fee = 4500
			default:
				fee = 5000
			}

			if subtotal >= FreeDeliveryThreshold {
				fee = 0
			}

			grandTotal := subtotal + fee
			fmt.Printf("%-12s | %8.2f ETB | %8.2f ETB | %8.2f ETB\n",
				area,
				float64(subtotal)/SantimPerBirr,
				float64(fee)/SantimPerBirr,
				float64(grandTotal)/SantimPerBirr,
			)
		}
	}
}
