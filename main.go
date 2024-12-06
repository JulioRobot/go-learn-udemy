package main

import (
	"fmt"
	"math"
)

func main() {
	calculate()
}

func invesmentRevenue() {
	// You can use like this
	// var investmentAmount , expectedReturnRate , years  float64 = 1000,5.5,10
	// investmentAmount , expectedReturnRate , years := 1000.0,5.5,10.0
	// expectedReturnRate := 5.5
	// years := 10.0

	const inflationRate = 6.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
