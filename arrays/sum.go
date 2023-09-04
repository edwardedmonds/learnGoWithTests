package main

import (
	"fmt"
)

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

func main() {
	sliceOne := []int{12, 23, 3, 4, 5}
	sliceTwo := []int{3, 5}
	sliceThree := []int{20, 39, 38}

	fmt.Println(SumAll(sliceOne, sliceTwo, sliceThree))
}
