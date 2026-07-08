package lists

import "fmt"

func main(){
	prices := []float64{10.99, 8.99}

	prices[1] = 9.99

	prices = append(prices, 5.99)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string= [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))

// }
