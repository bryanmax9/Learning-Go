package main

import (
	"fmt"
	"math"
)
const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Number of Years: ")
	outputText("Number of Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Println("The Future values is: ", futureValue)
	fmt.Println("The future Real values is: ",futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64){
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1 + inflationRate/100, years) 
	return fv, rfv 
}