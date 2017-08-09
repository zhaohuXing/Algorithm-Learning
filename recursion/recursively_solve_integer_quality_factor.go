package main

import (
	"fmt"
	"math"
)

// 试除法, try out the law
func isPrimeNumber(num int) bool {

	value := math.Sqrt(float64(num))
	for i := 2; i < (int)(value); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func maxPrimeNumber(num int) int {
	value := math.Sqrt(float64(num))
	for i := 2; i <= (int)(value); i++ {
		if num%i == 0 {
			return i
		}
	}
	return num
}

func F(num int, collection []int) []int {
	if isPrimeNumber(num) {
		return append(collection, num)
	} else {
		max := maxPrimeNumber(num)
		return F(num/max, append(collection, max))
	}
}

func main() {
	fmt.Println(F(2409, make([]int, 0)))
}
