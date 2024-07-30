package fazz

import (
	"fazztrack/fazzfood/calc"
	"fmt"
)

func FoodOrder(price int, disc string, distance int, tax bool) {
	userDisc := calc.Discount(price, disc)
	deliveryFee := calc.Distance(distance)
	taxValue := calc.Tax(price, tax)
	var total int

	total = price - userDisc + deliveryFee + taxValue
	fmt.Printf("Harga       : %d", price)
	fmt.Println("")
	fmt.Printf("Potongan    : %d", userDisc)
	fmt.Println("")
	fmt.Printf("Biaya Antar : %d", deliveryFee)
	fmt.Println("")
	fmt.Printf("Pajak       : %d", taxValue)
	fmt.Println("")
	fmt.Printf("SubTotal    : %d", total)
}
