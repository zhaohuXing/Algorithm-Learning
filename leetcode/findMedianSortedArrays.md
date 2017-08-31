# 4. Median of Two Sorted Arrays
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

##### Example 1:
```
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
```

#### Example 2:
```
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```

##### 解题思路：
  - 将两个有序数组看成两个列表
 - 之后, 变成将两个链表有序合并成一个链表

注: 这里重复的数据也需要加进去．

#### 实现代码：
```
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
```
