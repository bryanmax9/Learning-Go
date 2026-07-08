package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1
	hobbies := [3]string{"tennis", "basquetball", "golf"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//3

	highlightFirstSecond := hobbies[:2]

	fmt.Println(highlightFirstSecond)

	//4
	task4 := highlightFirstSecond[1:3]

	fmt.Println(task4)

	//5
	var dynamicArray = []string{"learn GO", "Master Go language", "Expert in Go"}
	fmt.Println(dynamicArray)

	//6
	dynamicArray[1] = "Create Go API"

	updatedDynamicArrray := append(dynamicArray, "Complete the Go course")

	fmt.Println(updatedDynamicArrray)

	//7
	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"second-product",
			"A Second Product",
			12.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"3",
		"Mew Toy",
		89.99,
	}

	products = append(products, newProduct)


	fmt.Println(products)


	//Unpacking List Values
	prices := []float64{10.99,8.99}
	
	prices = append(prices, 5.99, 12.99,29.99,100.10)

	fmt.Println(prices)

	discountPrices := []float64{101.99,80.99,20.59}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)

}
