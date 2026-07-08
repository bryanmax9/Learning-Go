package main

import "fmt"

func main(){
	userNames := make([]string, 2) //initial length of slice of 2

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	fmt.Println(userNames)

	courseRatings := make(map[string]float64,3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	fmt.Println(courseRatings)

	//For loops with slices
	for index,value := range userNames{
		fmt.Println(index,value)
	}
}
