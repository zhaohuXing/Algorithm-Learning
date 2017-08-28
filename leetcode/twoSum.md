# 1. Two Sum
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

#### Example:
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

解题思路：一边遍历数组时，借助 map 同时存储其值，这样扫一遍数组就可以得出结果，用空间换时间时间。(杭州面试题)

```
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
```


