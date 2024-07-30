package main

import (
	"fazztrack/fazzfood/fazz"
	"fmt"
)

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
	fazz.FoodOrder(orderPrice, voucher, distance, false)
}
