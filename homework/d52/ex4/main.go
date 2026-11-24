package main

import "fmt"

func main() {
	dishes := []string{
		"ዶሮ ወጥ",
		"ሽሮ ተጋቢኖ",
		"ክትፎ ልዩ ከአይብና ጎመን ጋር",
		"Special Beyaynetu",
		"Doro Wat with Injera & Ayib",
	}

	// len(s) returns the number of UTF-8 bytes in the string, whereas
	// len([]rune(s)) counts actual Unicode characters (runes). Each Amharic
	// character takes 3 bytes in UTF-8, while ASCII characters take 1 byte.
	for _, dish := range dishes {
		byteLen := len(dish)
		runeCount := len([]rune(dish))
		fits := runeCount <= 20

		fmt.Printf("Dish: %s\n", dish)
		fmt.Printf("  Bytes: %d | Runes: %d | Fits 20-char field: %v\n", byteLen, runeCount, fits)
	}
}
