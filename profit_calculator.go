package main

import (
	"fmt"
)

func calculate() {
	var revenue float64
	var expanses float64
	var taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("expanses: ")
	fmt.Scan(&expanses)

	fmt.Print("tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expanses
	println("EBT")
	println(ebt)

	profit := revenue - expanses - taxRate
	println("Profit")
	println(profit)

	ratio := ebt / profit
	println("Ratio")
	println(ratio)

}
