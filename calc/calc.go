package calc

func Discount(price int, disc string) int {
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
	}
	return userDisc
}

func Distance(distance int) int {
	var deliveryFee = 5000
	if deliveryFee > 2 {
		deliveryFee += (int(distance) - 2) * 3000
	}
	return deliveryFee
}

func Tax(price int, tax bool) int {
	var taxValue int
	if tax {
		taxValue = price * 5 / 100
	} else {
		taxValue = 0
	}
	return taxValue
}
