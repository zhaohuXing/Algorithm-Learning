#3. Longest Substring Without Repeating Characters

Given a string, find the length of the longest substring without repeating characters.

##### Examples:
Given `"abcabcbb"`, the answer is `"abc"`, which the length is 3.

Given `"bbbbb"`, the answer is `"b"`, with the length of 1.

Given `"pwwkew"`, the answer is `"wke"`, with the length of 3. Note that the answer must be a substring, `"pwke"` is a subsequence and not a substring.

#### 解题思路：
- 用两个 pointer 来标记最长子串的始终坐标，例如: left, right。right一直向后比遍历，通过查看 map 中的值，判断是否遇见重复的数据，left 则变为 `重复下标加一` 
- 利用 map 存放遍历过得 Key - Value
- 每次遍历计算出其长度，如果当前长度大于最长的长度，则更新最长的长度，否则不做改动。

#### 实现代码：
```
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	tbl := make(map[byte]int, n)
	var left, curr, longest int

	for right := 0; right < n; right++ {
		if k, ok := tbl[s[right]]; ok && k >= left {
			left = k + 1
		}

		tbl[s[right]] = right
		curr = right - left + 1

		if curr > longest {
			longest = curr
		}
	}
	return longest
}
```

