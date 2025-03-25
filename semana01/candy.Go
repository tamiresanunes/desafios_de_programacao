package main

import "fmt"

//Complexidade de tempo e espaço -> O(n)

func candy(ratings []int) int {

	candies := make([]int, len(ratings))

	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {

		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
		fmt.Println(candies)
	}

	for i := len(ratings) - 2; i >= 0; i-- {

		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}

		fmt.Println(candies)

	}
	fmt.Println("")

	fmt.Println(candies)

	totalCandies := 0
	for _, value := range candies {
		totalCandies += value
	}
	return totalCandies
}

func main() {
	children := []int{1, 2, 87, 87, 87, 2, 1}
	resultado := candy(children)
	fmt.Println(resultado)
}