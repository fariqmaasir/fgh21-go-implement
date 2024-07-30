package main

import (
	"fmt"
)

func fazzFood(price int, disc string, distance int, tax bool) {
	var userDisc int
	var taxValue int
	var total int
	// When user use FAZZFOOD50
	if disc == "FAZZFOOD50" && price >= 50000 {
		userDisc = price * 50 / 100
		if userDisc > 50000 {
			userDisc = 50000
		}
		// When user use DITRAKTIR60
	} else if disc == "DITRAKTIR60" && price >= 25000 {
		userDisc = price * 60 / 100
		if userDisc > 30000 {
			userDisc = 30000
		}
	} else if disc == "" {
		userDisc = 0
	} else {
		fmt.Println("YOUR DISCOUNT VOUCHER ISN'T VALID")
		return
	}

	// --------------Delivery Fee---------------
	var deliveryFee = 5000
	if deliveryFee > 2 {
		deliveryFee += (int(distance) - 2) * 3000
	}

	// --------------Tax---------------
	if tax {
		taxValue = price * 5 / 100
	} else {
		taxValue = 0
	}

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
	fazzFood(25000, "DITRAKTIR60", 5, true)
}
