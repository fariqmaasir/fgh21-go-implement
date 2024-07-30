package main

import (
	"fazztrack/fazzfood/calc"
	"fmt"
)

func fazzFood(price int, disc string, distance int, tax bool) {
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

func main() {
	var orderPrice int
	var voucher string
	var distance int
	fmt.Print("Masukan Harga  : ")
	fmt.Scanln(&orderPrice)
	fmt.Print("Masukan Diskon : ")
	fmt.Scanln(&voucher)
	fmt.Print("Masukan Jarak  : ")
	fmt.Scanln(&distance)
	fazzFood(orderPrice, voucher, distance, false)
}
