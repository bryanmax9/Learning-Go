package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5,1,2}

	triple := transformNumbers(&numbers, triple)

	fmt.Println(triple)

	tranformerFn1 := getTransformerFunction(&numbers)
	tranformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers,tranformerFn1)
	moreTransformedNumbers := transformNumbers(&numbers, tranformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn{
	if (*numbers)[0] == 1{
		return double
	} else{
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
