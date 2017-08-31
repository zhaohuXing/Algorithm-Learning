package main

import "fmt"

func main() {
	num := findMedianSortedArrays([]int{}, []int{2, 3})
	fmt.Println(num)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int
	nums3 := make([]int, 0)
	fmt.Println(len(nums1))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums3 = append(nums3, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			nums3 = append(nums3, nums2[j])
			j++
		} else if nums1[i] == nums2[j] {
			nums3 = append(nums3, nums1[i])
			nums3 = append(nums3, nums2[j])
			i++
			j++
		}
	}
	if j == len(nums2) {
		for i < len(nums1) {
			nums3 = append(nums3, nums1[i])
			i++
		}
	}
	if i == len(nums1) {
		for j < len(nums2) {
			nums3 = append(nums3, nums2[j])
			j++
		}
	}
	length := len(nums3)
	mid := length / 2
	if length%2 == 0 {
		return (((float64)(nums3[mid-1]) + (float64)(nums3[mid])) / 2)
	}
	return float64(nums3[mid])
}
