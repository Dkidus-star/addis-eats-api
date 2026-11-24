package main

import "fmt"

const SantimPerBirr = 100

type Dish struct {
	Name        string
	PriceSantim int
	IsFasting   bool
}

func main() {
	menu := []Dish{
		{Name: "ዶሮ ወጥ", PriceSantim: 45000, IsFasting: false},
		{Name: "ሽሮ ተጋቢኖ", PriceSantim: 18000, IsFasting: true},
		{Name: "ክትፎ ልዩ", PriceSantim: 52000, IsFasting: false},
		{Name: "በያይነቱ", PriceSantim: 22000, IsFasting: true},
		{Name: "Tibs Firfir", PriceSantim: 36000, IsFasting: false},
		{Name: "Misir Wat", PriceSantim: 16500, IsFasting: true},
		{Name: "Gomen Besiga", PriceSantim: 34000, IsFasting: false},
		{Name: "Timatim Kurt", PriceSantim: 14000, IsFasting: true},
	}

	if len(menu) == 0 {
		fmt.Println("Addis Eats Menu Report: menu is empty.")
		return
	}

	fmt.Println("=== Addis Eats Menu Report ===")

	cheapest := menu[0]
	mostExpensive := menu[0]
	totalSantim := 0
	fastingCount := 0
	nonFastingCount := 0

	for _, dish := range menu {
		var label string
		switch dish.IsFasting {
		case true:
			label = "Fasting"
			fastingCount++
		case false:
			label = "Non-fasting"
			nonFastingCount++
		}

		fmt.Printf("%-20s | %8.2f ETB | %s\n",
			dish.Name,
			float64(dish.PriceSantim)/SantimPerBirr,
			label,
		)

		totalSantim += dish.PriceSantim
		if dish.PriceSantim < cheapest.PriceSantim {
			cheapest = dish
		}
		if dish.PriceSantim > mostExpensive.PriceSantim {
			mostExpensive = dish
		}
	}

	// Integer division intentionally truncates the average to the santim.
	avgSantim := totalSantim / len(menu)

	fmt.Println("------------------------------")
	fmt.Printf("Cheapest dish:       %s (%.2f ETB)\n", cheapest.Name, float64(cheapest.PriceSantim)/SantimPerBirr)
	fmt.Printf("Most expensive dish: %s (%.2f ETB)\n", mostExpensive.Name, float64(mostExpensive.PriceSantim)/SantimPerBirr)
	fmt.Printf("Average price:       %.2f ETB\n", float64(avgSantim)/SantimPerBirr)
	fmt.Printf("Fasting dishes:      %d\n", fastingCount)
	fmt.Printf("Non-fasting dishes:  %d\n", nonFastingCount)
}
