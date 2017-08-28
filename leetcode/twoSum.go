package main

import (
	"fmt"
)

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{-1, -1}
	}

	other := make(map[int]int)

	for index, val := range nums {
		if i, ok := other[target-val]; ok {
			return []int{i, index}
		} else {
			other[val] = index
		}
	}
	return []int{-1, -1}
}
