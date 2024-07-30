package main

import (
	"fmt"
)

func calcDiscount(price int, disc string) int {
	var userDisc int
	if disc == "FAZZFOOD50" && price >= 50000 {
		userDisc = price * 50 / 100
		if userDisc > 50000 {
			userDisc = 50000
		}
		// When user use DITRAKTIR60
	} else if disc == "DITRAKTIR60" && price >= 25000 {
		userDisc = price * 60 / 100
		if userDisc >= 30000 {
			userDisc = 30000
		}
	} else if disc == "" {
		userDisc = 0
	} else {
		fmt.Println("YOUR DISCOUNT VOUCHER ISN'T VALID")
	}
	return userDisc
}

func calcDistance(distance int) int {
	var deliveryFee = 5000
	if deliveryFee > 2 {
		deliveryFee += (int(distance) - 2) * 3000
	}
	return deliveryFee
}

func calcTax(price int, tax bool) int {
	var taxValue int
	if tax {
		taxValue = price * 5 / 100
	} else {
		taxValue = 0
	}
	return taxValue
}
func fazzFood(price int, disc string, distance int, tax bool) {
	userDisc := calcDiscount(price, disc)
	deliveryFee := calcDistance(distance)
	taxValue := calcTax(price, tax)
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
