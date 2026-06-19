package main

import (
	"fmt"
	"errors"
	"os"
)


func main(){


	revenue, err1 := getUserInput("Input the Revenue amount: ")


	expenses, err2 := getUserInput("Input the Expenses amount: ")



	taxRate, err3 := getUserInput("Input the Tax Rate (%): %")

	if err1 != nil || err2 != nil || err3 != nil{
		fmt.Println(err1)
		return
	}

	EBT, profit, taxRatio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Your Earning Before Tax is: ", EBT)
	fmt.Println("Your Earning after Tax is: ", profit)
	fmt.Println("The Ratio result is: ", taxRatio)
	storeResults(EBT, profit, taxRatio)
}

func storeResults(ebt, profit,ratio float64){
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt",[]byte(results),0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64,float64){
	EBT := revenue - expenses

	profit := EBT * (1 - taxRate/100)

	taxRatio := EBT / profit

	return EBT,profit, taxRatio
}


func getUserInput(infoText string)  (float64, error){
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0,errors.New("Value must be a positve number")
	}
	return userInput, nil
}